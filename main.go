package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/julianossilva/chatappapi/handlers/hellohandler"
	"github.com/julianossilva/chatappapi/servicecontainer"
)

func main() {
	var err error
	app := fiber.New()

	var serviceContainer servicecontainer.ServiceContainer
	if serviceContainer, err = servicecontainer.New(); err != nil {
		panic(err.Error())
	}

	app.Use(servicecontainer.AppendServiceContainer(serviceContainer))

	app.Get("/", hellohandler.Hello())

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
