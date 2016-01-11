package manager

import (
	"testing"

	"github.com/seiflotfy/skizze/config"
	"github.com/seiflotfy/skizze/datamodel"
	"github.com/seiflotfy/skizze/storage"
	"github.com/seiflotfy/skizze/utils"
)

func TestInfoCreate(t *testing.T) {
	config.Reset()
	utils.SetupTests()
	defer utils.TearDownTests()

	s := storage.NewManager()
	m := newInfoManager(s)

	info := datamodel.NewEmptyInfo()
	info.Properties.Capacity = 10000
	info.Name = "marvel1"
	info.Type = datamodel.HLLPP

	if err := m.create(info); err != nil {
		t.Error("Expected no errors, got", err)
	}

}

func TestInfoSaveDelete(t *testing.T) {
	config.Reset()
	utils.SetupTests()
	defer utils.TearDownTests()

	s := storage.NewManager()
	m := newInfoManager(s)

	info := datamodel.NewEmptyInfo()
	info.Properties.Capacity = 10000
	info.Name = "marvel2"
	info.Type = datamodel.HLLPP

	if err := m.create(info); err != nil {
		t.Error("Expected no errors, got", err)
	}

	// Save state
	if err := m.save(info.ID()); err != nil {
		t.Error("Expected no errors, got", err)
	}

	// delete old Info
	if err := m.delete(info); err != nil {
		t.Error("Expected no errors, got", err)
	}

}

func TestInfoCreateDuplicate(t *testing.T) {
	config.Reset()
	utils.SetupTests()
	defer utils.TearDownTests()

	s := storage.NewManager()
	m := newInfoManager(s)

	info := datamodel.NewEmptyInfo()
	info.Properties.Capacity = 10000
	info.Name = "marvel3"
	info.Type = datamodel.HLLPP

	if err := m.create(info); err != nil {
		t.Error("Expected no errors, got", err)
	}

	// Save state
	if err := m.save(info.ID()); err != nil {
		t.Error("Expected no errors, got", err)
	}

	info2 := datamodel.NewEmptyInfo()
	info2.Properties.Capacity = 10000
	info2.Name = "marvel3"
	info2.Type = datamodel.HLLPP

	if err := m.create(info2); err == nil {
		t.Error("Expected errors, got", err)
	}

}

func TestInfoDeleteInvalid(t *testing.T) {
	config.Reset()
	utils.SetupTests()
	defer utils.TearDownTests()

	s := storage.NewManager()
	m := newInfoManager(s)

	info := datamodel.NewEmptyInfo()
	info.Properties.Capacity = 10000
	info.Name = "marvel4"
	info.Type = datamodel.HLLPP

	if err := m.delete(info); err == nil {
		t.Error("Expected errors, got", err)
	}

}