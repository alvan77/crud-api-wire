package mapper

import "crud-api-wire/models"

func ToProduct(productDTO models.ProductDTO) models.Product {
	return models.Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product models.Product) models.ProductDTO {
	return models.ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []models.Product) []models.ProductDTO {
	productdtos := make([]models.ProductDTO, len(products))
	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}
	return productdtos
}
