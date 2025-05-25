package main

import (
	"log"

	"github.com/abhishekjain1416/ad-campaign-service/api/route"
	"github.com/abhishekjain1416/ad-campaign-service/config"
	"github.com/abhishekjain1416/ad-campaign-service/db"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	db.InitRedis()

	r := gin.Default()
	route.RegisterRoutes(r)

	log.Println("Starting server on port 8080...")
	r.Run(":8080")
}
