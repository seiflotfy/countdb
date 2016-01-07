package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/seiflotfy/skizze/config"
	"github.com/seiflotfy/skizze/datamodel"
	"github.com/seiflotfy/skizze/utils"
)

// Manager the storage should deal with 2 types of on disk files, info and data
// info describes a domain and can be used to load back from disk the settings
// of a counter to reinitialize it
// the data is to refill the counters from disk
type Manager struct {
	db   *bolt.DB
	conf *config.Config
}

// NewManager ...
func NewManager() *Manager {
	conf := config.GetConfig()
	dataPath := conf.DataDir
	err := os.MkdirAll(dataPath, 0777)
	utils.PanicOnError(err)
	infoPath := filepath.Join(config.GetConfig().InfoDir, "info.db")
	db, err := bolt.Open(infoPath, 0777, nil)
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("sketches"))
		return err
	})
	utils.PanicOnError(err)
	return &Manager{db, conf}
}

// GetFile ...
func (m *Manager) GetFile(id string) (*os.File, error) {
	f, err := os.OpenFile(filepath.Join(m.conf.DataDir, id), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// SaveInfo ...
func (m *Manager) SaveInfo(info map[string]*datamodel.Info) error {
	return m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sketches"))
		for k, v := range info {
			rawInfo, err := json.Marshal(v)
			if err != nil {
				return fmt.Errorf("savin info: %s", err)
			}
			if err := b.Put([]byte(k), rawInfo); err != nil {
				return err
			}
		}
		return nil
	})
}

// LoadAllInfo ...
func (m *Manager) LoadAllInfo() (map[string]*datamodel.Info, error) {
	infos := map[string]*datamodel.Info{}
	err := m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sketches"))
		err := b.ForEach(func(k, v []byte) error {
			var info *datamodel.Info
			if err := json.Unmarshal(v, &info); err != nil {
				return err
			}
			infos[string(k)] = info
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return infos, nil
}

// Close ...
func (m *Manager) Close() error {
	return m.db.Close()
}
