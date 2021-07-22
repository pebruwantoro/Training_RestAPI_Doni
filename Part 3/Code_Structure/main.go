package main

import (
	"pebruwantoro/structuring/config"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
