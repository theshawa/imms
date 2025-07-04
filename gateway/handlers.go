package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/logan2k02/ims/gateway/inventory_handlers"
	"github.com/logan2k02/ims/gateway/orders_handlers"
	"github.com/logan2k02/ims/gateway/products_handlers"
	pb "github.com/logan2k02/ims/shared/protobuf"

	_ "github.com/joho/godotenv/autoload"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func registerHandlers(app *fiber.App, productsClient pb.ProductsServiceClient, inventoryClient pb.InventoryServiceClient, ordersClient pb.OrdersServiceClient) {
	app.Post("/products/create", products_handlers.CreateProductHandler(productsClient, validate))
	app.Get("/products/:id", products_handlers.GetProduct(productsClient))
	app.Get("/products", products_handlers.ListProducts(productsClient))
	app.Delete("/products/:id", products_handlers.DeleteProduct(productsClient))
	app.Put("/products/:id", products_handlers.UpdateProduct(productsClient, validate))

	app.Post("/inventory/supply/:id", inventory_handlers.Supply(inventoryClient, validate))
	app.Post("/inventory/correct/:id", inventory_handlers.Correct(inventoryClient, validate))

	app.Post("/orders/create", orders_handlers.CreateOrderHandler(ordersClient, validate))
	app.Get("/orders", orders_handlers.ListOrdersHandler(ordersClient))
	app.Get("/orders/:id", orders_handlers.GetOrderHandler(ordersClient))
	app.Post("/orders/change-status/:id", orders_handlers.ChangeOrderStatusHandler(ordersClient, validate))
	app.Delete("/orders/:id", orders_handlers.DeleteOrderHandler(ordersClient))
}
