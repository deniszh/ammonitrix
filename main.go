package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
	"github.com/eBayClassifiedsGroup/ammonitrix/receiver"
)

var version = "0.0.1-dev"

func main() {
	var filename string
	var v bool
	flag.StringVar(&filename, "cfg", "", "path to config file")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		fmt.Println(version)
		return
	}
	log.Printf("[INFO] Version %s starting", version)

	cfg := config.DefaultConfig
	if filename != "" {
		cfg = loadConfig(filename)
	}

	recv, err := receiver.NewReceiver(cfg)
	if err != nil {
		log.Fatal("[FATAL] Unable to start Receiver:", err)
	}

	metadata, err := recv.ConnectElastic()
	if err != nil {
		log.Fatal("[FATAL] Unable to connect to elasticsearch:", err)
	}

	err = recv.StartListener(metadata)
	if err != nil {
		log.Fatal("[FATAL] Unable to start listener:", err)
	}

}
