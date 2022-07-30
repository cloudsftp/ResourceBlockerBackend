package main

import (
	"flag"

	"github.com/cloudsftp/ResourceBlockerBackend/server"
)

func main() {
	pconfig := flag.String("config", "config.json", "configuration file")
	flag.Parse()

	config := server.GetConfig(*pconfig)

	server.StartServer(config)
}
