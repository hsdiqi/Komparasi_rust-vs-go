package main

import (
	"log"

	"backend-golang-gin/config"
	"backend-golang-gin/database"
	"backend-golang-gin/handler"
	"backend-golang-gin/repository"
	"backend-golang-gin/route"
	"backend-golang-gin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db, err := database.NewPostgres(cfg)
	if err != nil { log.Fatal(err) }
	defer db.Close()

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	h := handler.NewProductHandler(svc)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	route.Register(r, h)

	log.Println("Golang Gin running on :" + cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil { log.Fatal(err) }
}
