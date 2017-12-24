package main

import (
	"fmt"
	"log"

	"github.com/turtledev/refract-api/web"
)

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	server := web.NewServer(addr)
	log.Println("Refract API Server", version)
	log.Println("starting on", addr)
	log.Fatalln(server.Start())
}
