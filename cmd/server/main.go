package main

import (
	"fmt"
	"net/http"

	"github.com/matheusteixeira7/fc-go-api/configs"
	"github.com/matheusteixeira7/fc-go-api/internal/entity"
	"github.com/matheusteixeira7/fc-go-api/internal/infra/database"
	"github.com/matheusteixeira7/fc-go-api/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
