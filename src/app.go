package main

import (
	post "actions/post"
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
	fw.Get("/post/:id", post.Read)

	<-ctx.Done()
}
