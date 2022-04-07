package main

import (
	"crud-api-wire/models"
	"crud-api-wire/wire"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	dbUser := "root"
	log.Println(dbUser)
	pass := "root"
	host := "localhost"
	port := "3306"
	dbName := "api_db"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, pass, host, port, dbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Product{})
	return db
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := wire.InitProductAPI(db)

	r := gin.Default()

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
