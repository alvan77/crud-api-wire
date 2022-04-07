package service

import (
	"crud-api-wire/models"
	"crud-api-wire/repository"
)

// mendefinisikan struct product service yg membawa product repository dari folder repository

type ProductService interface {
	FindAll() []models.ProductDTO
}
type productService struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductService(p repository.ProductRepository) ProductService {
	return &productService{ProductRepository: p}
}

func (p *productService) FindAll() []models.ProductDTO {
	return p.ProductRepository.FindAll()
}

// func (p *ProductService) FindByID(id uint) models.Product {
// 	return p.ProductRepository.FindByID(id)
// }

// func (p *ProductService) Save(product models.Product) models.Product {
// 	p.ProductRepository.Save(product)
// 	return product
// }

// func (p *ProductService) Delete(product models.Product) {
// 	p.ProductRepository.Delete(product)
// }
