package main

import (
	"CrowdProject/internal/config"
	"CrowdProject/server"
	"fmt"
	"log"
)

func main() {
	cfg := config.GetConfig()

	fmt.Println(cfg.Auth.TokenTtl)

	app := server.NewApp(cfg)

	err := app.Run("8000")
	if err != nil {
		log.Fatalf("The app can't run, err =%s", err)
	}

}
