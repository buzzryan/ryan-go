package main

import "github.com/buzzryan/ryan-go/api"

func main() {
	initialize(nil)
	launchServer()
}

func initialize(c *config) {

}

func launchServer() {
	app := api.InitializeServerApp()
	app.Run()
}

type config struct {
}
