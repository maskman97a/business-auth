package main

import "business-auth/conf"

func main() {
	app := conf.NewApp()
	app.Init()
	app.Run()
}
