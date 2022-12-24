package main

import (
	"log"

	"example/pubg-stats-api/client"
)

func main() {
	cl, err := client.New()
	if err != nil {
		log.Panicln(err)
	}
	runServer(cl)
}
