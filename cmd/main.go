package main

import (
	"log"

	"github.com/bintangyosua/echaaas/bootstrap"
	"github.com/bintangyosua/echaaas/configs"
	"github.com/bintangyosua/echaaas/internal/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system env vars")
	}

	db := configs.InitDB()
	container := bootstrap.BuildContainer(db)

	auth.Module{}.RegisterRoutes(e, container)


	e.Logger.Fatal(e.Start(":8080"))
}