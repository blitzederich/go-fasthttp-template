# Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

BUILD_DIR = $(CURDIR)/bin
PKG_NAME = api

clean:
	rm -rf $(CURDIR)/bin

build:
	go build -o $(BUILD_DIR)/$(PKG_NAME) $(CURDIR)/cmd/$(PKG_NAME)/main.go
	cp ./config.json $(BUILD_DIR)/

service_create:
	sudo cp ./$(PKG_NAME).service /usr/lib/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl start $(PKG_NAME).service
	sudo systemctl status $(PKG_NAME).service

service_remove:
	sudo systemctl stop $(PKG_NAME).service
	sudo rm /usr/lib/systemd/system/$(PKG_NAME).service
	sudo systemctl daemon-reload