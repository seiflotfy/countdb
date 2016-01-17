package main

import (
	"flag"
	"os"
	"strconv"

	_ "net/http/pprof"

	"github.com/skizzehq/skizze/config"
	"github.com/skizzehq/skizze/manager"
	"github.com/skizzehq/skizze/server"
	"github.com/skizzehq/skizze/utils"
)

var logger = utils.GetLogger()

func main() {
	var port uint
	flag.UintVar(&port, "p", 3596, "specifies the port for Skizze to run on")
	flag.Parse()

	//TODO: Add arguments for dataDir and infoDir

	err := os.Setenv("SKIZZE_PORT", strconv.Itoa(int(port)))
	utils.PanicOnError(err)

	logger.Info.Println("Starting Skizze...")
	logger.Info.Println("Using data dir: ", config.GetConfig().DataDir)
	//server, err := server.New()
	//utils.PanicOnError(err)
	//server.Run()
	mngr := manager.NewManager()
	if p, err := strconv.Atoi(os.Getenv("SKIZZE_PORT")); err == nil {
		server.Run(mngr, uint(p))
	}
}
