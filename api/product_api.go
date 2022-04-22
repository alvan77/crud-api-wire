package api

import (
	"crud-api-wire/models"
	"crud-api-wire/service"
	"crud-api-wire/utils"
	"log"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ProductAPI struct {
	ProductService service.ProductService
}

func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	name := c.Query("name")
	var dateStartTime, dateEndTime time.Time
	var err error
	if startDate != "" {
		dateStartTime, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	if endDate != "" {
		dateEndTime, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	log.Println(startDate, endDate, "date from api")

	if name != "" {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	log.Println(name, "name from api")

	products, err := p.ProductService.FindAll(dateStartTime, dateEndTime, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// func (p *ProductAPI) FindByName(c *gin.Context) {
// 	name := c.Query("name")
// 	var err error
// 	if name != "" {
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 	}

// 	log.Println(name, "name api")

// 	products, err := p.ProductService.FindByName(string(name))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"products": products})
// }

func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"products": product})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"errors": err.Error()})

		return
	}
	validate := validator.New()
	utils.GetJsonTag(validate)
	err := validate.Struct(product)
	if err != nil {
		errMessage := utils.ErrorValidationMessage(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": errMessage})
		return
	}

	createdProduct := p.ProductService.Save(product)
	c.JSON(http.StatusOK, gin.H{"product": createdProduct})
}

// func (p *ProductAPI) Update(c *gin.Context) {
// 	var products models.Product
// 	err := c.BindJSON(&products)
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
// 	products.ID = product.ID
// 	products.Code = product.Code
// 	products.Price = product.Price
// 	p.ProductService.Save(products)
// 	c.Status(http.StatusOK)
// }

// func (p *ProductAPI) Delete(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	product := p.ProductService.FindByID(uint(id))
// 	if product == (models.Product{}) {
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := p.ProductService.Delete(product)
// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }
