package main

import (
	"log"

	"github.com/turtledev/refract-api/web"
)

func main() {
	server := web.NewServer(":8080")
	log.Println("starting api server on port 8080")
	log.Fatalln(server.Start())
}
