package main

import (
	"github.com/turtledev/refract-api/web"
)

func main() {
	server := web.NewServer()
	server.Start()
}
