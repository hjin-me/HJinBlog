package main

import (
	"framework"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered in main %v\n", r)
		}
	}()

	ctx := fw.App()
	// route init
	fw.Get("/a", func(ctx fw.Context) {
	})

	<-ctx.Done()
}
