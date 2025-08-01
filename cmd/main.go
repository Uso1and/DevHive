package main

import (
	"devhive/internal/app/handlers"
	"devhive/internal/domain/indrastructure/database"
	"devhive/internal/domain/repo"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	userRepo := repo.NewUserRepo(database.DB)
	userHandler := handlers.NewUserHandler(userRepo)

	r := gin.Default()

	r.LoadHTMLGlob("template/*")
	r.Static("static", "./static")

	r.POST("/signup", userHandler.SingnUpHandler)
	r.POST("/login", userHandler.LoginHandler)

	r.GET("/", handlers.IndexPageHandler)

	r.GET("/signup", handlers.SignUpPageHandler)
	r.GET("/profile", handlers.ProfilePageHandler)

	r.Run(":8080")

}
