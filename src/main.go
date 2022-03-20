package main

import (
	"github.com/hermesmartins/golang-fiber-abstraction/src/infra/http/fiber"
	"github.com/hermesmartins/golang-fiber-abstraction/src/infra/http/router"
)

func main() {
	http := fiber.NewFiber()
	router.NewRouter(http)
	http.Listen(":8080")
}
