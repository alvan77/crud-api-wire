package main

import (
	"crud-api-wire/models"
	"crud-api-wire/wire"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{})

	return db
}

func main() {
	db := initDB()
	defer db.Close() // fungsi defer ?

	productAPI := wire.InitProductAPI(db)

	r := gin.Default() // kenapa harus diseting default ?

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/detail/:id", productAPI.FindByID)
	r.POST("/products/add", productAPI.Create)
	r.PUT("/products/update/:id", productAPI.Update)
	r.DELETE("/products/delete/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
