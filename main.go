package main

import (
	"log"
	"notifs/src/config"
	"notifs/src/db"
	"notifs/src/server"
	"notifs/src/types"
)

var cfg *types.Config

func init() {
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
