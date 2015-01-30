package main

import (
	"actions"
	"framework"
	"os"
)

func main() {
	fw.Init()
	fw.Post("/cp", actions.Create)
	fw.Get("/cp", actions.ReadBatch)
	fw.Get("/cp/:id", actions.Read)
	fw.Put("/cp/:id", actions.Update)
	fw.Delete("/cp/:id", actions.Delete)
	fw.Get("/install", actions.LevelInit)
	fw.Get("/:prefix/:id", actions.PostRead)
	fw.Get("/", actions.PostReadBatch)
	fw.File("/statics", os.Getenv("GOPATH")+"/src/statics")
	// fw.Get("/:abc/:er", sayHello)
	fw.App()
}
