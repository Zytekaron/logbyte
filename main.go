package main

import (
	"log"
	"logbyte/src/config"
	"logbyte/src/db"
	"logbyte/src/server"
	"logbyte/src/types"
	"os"
)

var cfg *types.Config

func init() {
	if os.Getenv("ZYTEKARON_AUTH") == "" {
		log.Fatal("ZYTEKARON_AUTH is empty")
	}

	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Init(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server.Start(cfg.Server)
}
