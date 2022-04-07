package mapper

import "crud-api-wire/models"

func ToProduct(productDTO models.ProductDTO) models.Product {
	return models.Product{AssingedTo: productDTO.AssingedTo, Task: productDTO.Task, Deadline: productDTO.Deadline}
}

func ToProductDTO(product models.Product) models.ProductDTO {
	return models.ProductDTO{ID: product.ID, AssingedTo: product.AssingedTo, Task: product.Task, Deadline: product.Deadline}
}

func ToProductDTOs(products []models.Product) []models.ProductDTO {
	productdtos := make([]models.ProductDTO, len(products))
	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}
	return productdtos
}
