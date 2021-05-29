package repository

import (
	"github.com/donreno/gofiber-test-api/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	Get(ID uint) (model.Product, error)
	Create(model.Product) (model.Product, error)
	Update(model.Product) (model.Product, error)
	Delete(model.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func MakeProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAll() ([]model.Product, error) {
	products := []model.Product{}

	result := r.db.Find(&products)

	return products, result.Error
}

func (r *productRepository) Get(ID uint) (model.Product, error) {
	product := new(model.Product)

	result := r.db.Find(product, ID)

	return *product, result.Error
}

func (r *productRepository) Create(product model.Product) (model.Product, error) {
	result := r.db.Create(&product)

	return product, result.Error
}

func (r *productRepository) Update(product model.Product) (model.Product, error) {
	result := r.db.Save(&product)

	return product, result.Error
}

func (r *productRepository) Delete(product model.Product) error {
	result := r.db.Delete(&product)

	return result.Error
}
