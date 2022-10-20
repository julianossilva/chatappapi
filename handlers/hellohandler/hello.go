package hellohandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julianossilva/chatappapi/servicecontainer"
)

func Hello() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error

		var serviceContainer servicecontainer.ServiceContainer
		if serviceContainer, err = servicecontainer.GetServiceContainer(c); err != nil {
			return err
		}

		return c.SendString(serviceContainer.GetMessage())
	}
}
