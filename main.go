package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"metodichka-backend/db_config"
	"metodichka-backend/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return
	}

	db := db_config.ConnectDB()
	if err := db_config.Migrate(db); err != nil {
		log.Fatal(err)
		return
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	def := r.Group("api")
	def.GET("categories", handlers.GetCategories(db))
	def.GET("news", handlers.GetNews(db))
	def.POST("news", handlers.SetNews(db))

	def.GET("pages/:id", handlers.GetPage(db))

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

}
