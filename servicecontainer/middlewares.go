package servicecontainer

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func AppendServiceContainer(sc ServiceContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {

		c.Locals("servicecontainer", sc)

		return c.Next()
	}
}

func GetServiceContainer(c *fiber.Ctx) (ServiceContainer, error) {
	serviceContainer := c.Locals("servicecontainer")

	if serviceContainer == nil {
		return nil, errors.New("service container not exist")
	}

	if serviceContainer, ok := serviceContainer.(ServiceContainer); ok {
		return serviceContainer, nil
	} else {
		return nil, errors.New("service container not exist")
	}
}
