package repository

import (
	"crud-api-wire/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// mendefinisikan struct product repository yg berisi database dan membawa package gorm

type ProductRepository interface {
	FindAll() []models.ProductDTO
}
type productRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) FindAll() []models.ProductDTO {
	var products []models.ProductDTO
	var err error
	log.Println("error")
	err = p.DB.Find(&products).Error
	if err != nil {
		log.Println(err)
	}
	return products
}

func (p *productRepository) FindByID(id uint) models.Product {
	var product models.Product
	p.DB.First(&product, id)
	return product
}

func (p *productRepository) Save(product models.Product) models.Product {
	p.DB.Save(&product)

	return product
}

func (p *productRepository) Delete(product models.Product) {
	p.DB.Delete(&product)
}
