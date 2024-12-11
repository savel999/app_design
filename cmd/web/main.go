package main

import (
	"log"

	_ "time/tzdata"

	"github.com/joho/godotenv"
	"github.com/savel999/app_design/internal/app/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	config := web.InitConfig()

	if err := web.Run(config); err != nil {
		log.Fatal("failed to run app", err)
	}
}
