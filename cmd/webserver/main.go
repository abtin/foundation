package main

import (
	"log"

	"foundation/pkg/app"
	"foundation/pkg/config"
)

func main() {
	cfg, err := config.FromFile("./dev-config.toml")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	server := app.NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatalf("%s\n", err)
	}
}
