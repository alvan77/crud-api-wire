package repository

import (
	"crud-api-wire/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// mendefinisikan struct product repository yg berisi database dan membawa package gorm
type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []models.Product {
	var products []models.Product
	p.DB.Find(&products)
	return products
}

func (p *ProductRepository) FindByID(id uint) models.Product {
	var product models.Product
	p.DB.First(&product, id)
	return product
}

func (p *ProductRepository) Save(product models.Product) models.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product models.Product) {
	p.DB.Delete(&product)
}
