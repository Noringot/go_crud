package endpoint

import (
	"fmt"
	"github.com/Noringotq/go-crud/internal/app/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Initialize() {
	app := fiber.New()
	route.SetupRoute(app)

	err := app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server started")
}
