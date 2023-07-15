package main

import (
	"log"
	"menuProvider/database"
	"menuProvider/handler"
	"menuProvider/repository"
	"menuProvider/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()

	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	router := gin.Default()

	router.GET("/menu", h.Get)
	if err := router.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
