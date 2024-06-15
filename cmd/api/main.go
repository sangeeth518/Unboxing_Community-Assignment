package main

import (
	"Assignment/pkg/config"
	"Assignment/pkg/db"
	"Assignment/pkg/di"
	"fmt"
	"log"
)

func main() {
	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config", configerr)
	}
	fmt.Println(config.DBName)
	db.ConnectDB(config)

	server, dierr := di.InitializeAPI(config)

	if dierr != nil {
		log.Fatal("cannot start server", dierr)
	} else {
		server.Start()
	}

}
