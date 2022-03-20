package router

import (
	"github.com/hermesmartins/golang-fiber-abstraction/src/application/usecases"
	"github.com/hermesmartins/golang-fiber-abstraction/src/infra/http/fiber"
)

type Router struct {
	http *fiber.Fiber
}

func NewRouter(http *fiber.Fiber) *Router {
	router := &Router{
		http: http,
	}
	router.Init()
	return router
}

func (router Router) Init() {
	router.http.Get("/commands/cmd-1", func(params []string, body []byte) string {
		cmd1 := usecases.NewCmd1()
		return cmd1.Execute()
	})
	router.http.Get("/commands/cmd-2", func(params []string, body []byte) string {
		cmd2 := usecases.NewCmd2()
		return cmd2.Execute()
	})
}
