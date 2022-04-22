package main

import (
	// "crud-api-wire/db_app"
	"crud-api-wire/config"
	"crud-api-wire/models"
	"fmt"

	// "crud-api-wire/config"
	// "crud-api-wire/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func initDB() *gorm.DB {
	DBUser := viper.GetString("DB_USER")
	DBPass := viper.GetString("DB_PASS")
	DBHost := viper.GetString("DB_HOST")
	DBPort := viper.GetString("DB_PORT")
	DBName := viper.GetString("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPass, DBHost, DBPort, DBName)

	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Product{})
	return db
}

func main() {
	config.LoadConfig()
	db := initDB()
	defer db.Close()

	productAPI := ProductHandler(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	// r.GET("/products/detail", productAPI.FindByName)
	r.GET("/products/detail/:id", productAPI.FindByID)
	r.POST("/products/add", productAPI.Create)
	// r.PUT("/products/update/:id", productAPI.Update)
	// r.DELETE("/products/delete/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
