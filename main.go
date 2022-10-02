package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	//user
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("api/users/:id", routes.GetUser)
	app.Put("api/users/:id", routes.UpdateUser)
	app.Delete("api/users/:id", routes.DeleteUser)
	//products
	app.Post("api/products", routes.CreateProduct)
	app.Get("api/products", routes.GetProducts)
	app.Get("api/products/:id", routes.GetProduct)
	app.Put("api/products/:id", routes.UpdateProduct)
	app.Delete("api/products/:id", routes.DeleteProduct)
	//order
	app.Post("api/order", routes.CreateOrder)
	app.Get("api/order", routes.GetOrders)
	app.Get("api/order/:id", routes.GetOrder)

}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
