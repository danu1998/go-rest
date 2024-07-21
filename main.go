package main

import (
	productcontroller "github.com/danu1998/go-rest/controllers"
	"github.com/danu1998/go-rest/models"
	"github.com/gin-gonic/gin"
)

// const DB_USERNAME = "root"
// const DB_PASSWORD = ""
// const DB_NAME = "go_database"
// const DB_HOST = "127.0.0.1"
// const DB_PORT = "3306"

// dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
// fmt.Println("dsn : ", dsn)
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// if err != nil {
// 	panic("Unable Connect")
// }

// fmt.Println(db)
func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
