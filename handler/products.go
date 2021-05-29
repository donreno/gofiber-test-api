package handler

import (
	"fmt"

	"github.com/donreno/gofiber-test-api/model"
	"github.com/donreno/gofiber-test-api/repository"
	"github.com/gofiber/fiber/v2"
)

type ProductsHandler struct {
	repo repository.ProductRepository
}

func MakeProductsHandler(repo repository.ProductRepository) *ProductsHandler {
	return &ProductsHandler{
		repo: repo,
	}
}

func (p *ProductsHandler) GetAll(c *fiber.Ctx) error {
	products, err := p.repo.GetAll()

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if len(products) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"error": "Products not found",
		})
	}

	return c.Status(200).JSON(&products)
}

func (p *ProductsHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": fmt.Sprintf("error invalid id %q param, value should be numeric", c.Params("id")),
		})
	}

	product, err := p.repo.Get(uint(id))

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if product.ID == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"error": "Product not found",
		})
	}

	return c.Status(200).JSON(&product)
}

func (p *ProductsHandler) Update(c *fiber.Ctx) (err error) {
	product := model.Product{}

	if err = c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": fmt.Sprintf("bad request body %s", err.Error()),
		})
	}

	if product, err = p.repo.Update(product); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": fmt.Sprintf("error updating product %s", err.Error()),
		})
	}

	return c.Status(200).JSON(&product)
}

func (p *ProductsHandler) Create(c *fiber.Ctx) (err error) {
	product := model.Product{}

	if err = c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": fmt.Sprintf("bad request body %s", err.Error()),
		})
	}

	if product, err = p.repo.Create(product); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": fmt.Sprintf("error creating product %s", err.Error()),
		})
	}

	return c.Status(200).JSON(&product)
}

func (p *ProductsHandler) Delete(c *fiber.Ctx) (err error) {
	product := model.Product{}

	if err = c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": fmt.Sprintf("bad request body %s", err.Error()),
		})
	}

	if err = p.repo.Delete(product); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": fmt.Sprintf("error updating product %s", err.Error()),
		})
	}

	return c.SendStatus(200)
}
