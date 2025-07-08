package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/skncvo/Scalable_API/app/router"
	"github.com/skncvo/Scalable_API/config"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
