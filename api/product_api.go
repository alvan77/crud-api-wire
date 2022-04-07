package api

// import file yg dibutuhkan
import (
	"crud-api-wire/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// mendefinisikan struct product api yg membawa product service dari folder service
type ProductAPI struct {
	ProductService service.ProductService
}

// mendefinisikan function provide product api dengan parameter p yg berisi product service
// dari folder service dan memiliki nilai kembalian product api
func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

// mendefinisikan recievier
func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()

	// c.JSON(http.StatusOK, gin.H{"products": mapper.ToProductDTOs(products)})
	c.JSON(http.StatusOK, gin.H{"products": products})

}

// mendefinisikan reciever dengan parameter p, dan membuat method findByID dengan parameter c
// func (p *ProductAPI) FindByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	product := p.ProductService.FindByID(uint(id))

// 	c.JSON(http.StatusOK, gin.H{"product": mapper.ToProductDTO(product)})
// }

// func (p *ProductAPI) Create(c *gin.Context) {
// 	var productDTO models.ProductDTO
// 	err := c.BindJSON(&productDTO)
// 	if err != nil {
// 		log.Fatalln(err)
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}

// 	createdProduct := p.ProductService.Save(mapper.ToProduct(productDTO))
// 	c.JSON(http.StatusOK, gin.H{"product": mapper.ToProductDTO(createdProduct)})
// }

// func (p *ProductAPI) Update(c *gin.Context) {
// 	var productDTO models.ProductDTO
// 	err := c.BindJSON(&productDTO)
// 	if err != nil {
// 		log.Fatalln(err)
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	product := p.ProductService.FindByID(uint(id))
// 	if product == (models.Product{}) {
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}
// 	product.AssingedTo = productDTO.AssingedTo
// 	product.Task = productDTO.Task
// 	product.Deadline = productDTO.Deadline
// 	p.ProductService.Save(product)

// 	c.Status(http.StatusOK)
// }

// func (p *ProductAPI) Delete(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	product := p.ProductService.FindByID(uint(id))
// 	if product == (models.Product{}) {
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}

// 	p.ProductService.Delete(product)

// 	c.Status(http.StatusOK)
// }
