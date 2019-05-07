package main

import (
	"fmt"

	"./app"
	"./config"
)

func main() {
	config := config.GetConfig()
	fmt.Println("on progress")

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
