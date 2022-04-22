package service

import (
	"crud-api-wire/models"
	"crud-api-wire/repository"
	"time"
)

// mendefinisikan struct product service yg membawa product repository dari folder repository

type ProductService interface {
	FindAll(startDate time.Time, endDate time.Time, name string) ([]models.Product, error)
	// FindByName(name string) ([]models.Product, error)
	FindByID(id uint) models.Product
	Save(product models.Product) models.Product
	Delete(product models.Product) (models.Product, error)
}

type productService struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductService(p repository.ProductRepository) ProductService {
	return &productService{ProductRepository: p}
}

func (p *productService) FindAll(startDate time.Time, endDate time.Time, name string) ([]models.Product, error) {
	return p.ProductRepository.FindAll(startDate, endDate, name)
}

// func (p *productService) FindByName(name string) ([]models.Product, error) {
// 	return p.ProductRepository.FindByName(name)
// }

func (p *productService) FindByID(id uint) models.Product {
	return p.ProductRepository.FindByID(id)
}

func (p *productService) Save(product models.Product) models.Product {
	p.ProductRepository.Save(product)
	return product
}

func (p *productService) Delete(product models.Product) (models.Product, error) {
	data, err := p.ProductRepository.Delete(product.ID)
	if err != nil {
		return models.Product{}, err
	}
	return data, nil
}
