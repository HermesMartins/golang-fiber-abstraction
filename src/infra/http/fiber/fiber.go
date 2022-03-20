package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hermesmartins/golang-fiber-abstraction/src/infra/http/contracts"
)

type Fiber struct {
	app *fiber.App
}

func NewFiber() *Fiber {
	return &Fiber{
		app: fiber.New(),
	}
}

func (f Fiber) Get(url string, fn contracts.Fn) {
	f.app.Get(url, func(c *fiber.Ctx) error {
		response := fn(c.Route().Params, c.Body())
		return c.JSON(response)
	})
}

func (f Fiber) Listen(port string) {
	f.app.Listen(port)
}
