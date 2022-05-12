// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	config *Config
)

type Config struct {
	ServerAddr string `json:"server_addr"`
	JwtSecret  string `json:"jwt_secret"`
	Postgres   struct {
		ConnStr string `json:"conn_str"`
	} `json:"postgres"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
}

func Setup(configPath, serverAddr string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	config = new(Config)
	if err := json.Unmarshal(data, config); err != nil {
		return err
	}

	if serverAddr != "" {
		config.ServerAddr = serverAddr
	}

	return nil
}

func GetConfig() *Config {
	if config == nil {
		log.Fatal("Config was not initialized. Use config.Setup(configPath) for init.")
		os.Exit(1)
	}
	return config
}

func GetConfigCopy() Config {
	return *config
}
