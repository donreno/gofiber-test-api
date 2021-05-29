package wire

import (
	"fmt"

	"github.com/donreno/gofiber-test-api/database"
	"github.com/donreno/gofiber-test-api/handler"
	"github.com/donreno/gofiber-test-api/repository"
)

// Wires all dependencies to be injected on handler

func MakeHandler() *handler.ProductsHandler {
	db, err := database.Connect()
	if err != nil {
		panic(fmt.Sprintf("error loading connection: %s", err.Error()))
	}

	database.Migrate(db)
	if err != nil {
		panic(fmt.Sprintf("error migrating DB: %s", err.Error()))
	}

	repo := repository.MakeProductRepository(db)

	return handler.MakeProductsHandler(repo)
}
