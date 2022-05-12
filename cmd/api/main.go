// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package main

import (
	"api/internal/config"
	"api/internal/env"
	"api/internal/handlers"
	"api/internal/server"
	"flag"
	"log"
	"os"
)

var (
	argv struct {
		configPath string
		serverAddr string
	}
)

func main() {
	flag.StringVar(&argv.configPath, "config_path", "./config.json", "the path to the configuration file")
	flag.StringVar(&argv.serverAddr, "server_addr", "", "server address, for example :3000 ")
	flag.Parse()

	if err := config.Setup(argv.configPath, argv.serverAddr); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	config := config.GetConfig()
	if err := env.Setup(config); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	s := server.New()

	s.POST("/login", handlers.Login)

	if err := s.Start(config.ServerAddr); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
