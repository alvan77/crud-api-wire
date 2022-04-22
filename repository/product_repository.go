package repository

import (
	"crud-api-wire/models"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepository interface {
	FindAll(startDate time.Time, endDate time.Time, name string) ([]models.Product, error)
	// FindByName(name string) ([]models.Product, error)
	FindByID(id uint) models.Product
	Save(product models.Product) models.Product
	// Delete(product models.Product)
	Delete(id uint) (models.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) FindAll(startDate time.Time, endDate time.Time, name string) ([]models.Product, error) {
	var products []models.Product
	var err error

	query := p.DB.Debug()
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at  BETWEEN ? AND ? ", startDate, endDate)
	}

	if name != "" {
		query = query.Where("code LIKE ?", "%"+name+"%")

	}

	err = query.Find(&products).Error
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

// func (p *productRepository) FindByName(name string) ([]models.Product, error) {
// 	var products []models.Product
// 	if name != "" {
// 		log.Println(name)
// 	}

// 	err := p.DB.Debug().Where("code LIKE ?", "%"+name+"%").Find(&products)
// 	if err.Error != nil {
// 		return nil, err.Error
// 	}

// 	return products, nil
// }

func (p *productRepository) FindByID(id uint) models.Product {
	var product models.Product
	p.DB.First(&product, id)

	return product
}

func (p *productRepository) Save(product models.Product) models.Product {
	var err error
	log.Println(product)
	err = p.DB.Save(&product).Error
	if err != nil {
		log.Println(err)
	}
	return product
}

func (p *productRepository) Delete(id uint) (models.Product, error) {
	var product models.Product
	var err error
	log.Println(product)
	err = p.DB.Delete(&product, id).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
