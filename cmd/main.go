package main

import "business-auth/config"

func main() {
	app := config.NewApp()
	app.Init()
	app.Run()
}
