package main

import (
	"log"

	"github.com/Falchizao/api-golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro porra")
	}

	router := gin.Default()
	routes.Init(&router.RouterGroup)
	if err := router.Run(":8083"); err != nil {
		log.Fatal(err)
	}
}
