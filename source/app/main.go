package main

import (
	"log"
	"todo/app/presentation/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()
	router.IncludeRouter(e)
	e.Logger.Fatal(e.Start(":8000"))
}
