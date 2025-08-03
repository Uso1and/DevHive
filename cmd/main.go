package main

import (
	"devhive/internal/app/handlers"
	"devhive/internal/app/middleware"
	"devhive/internal/domain/config"
	"devhive/internal/domain/indrastructure/database"
	"devhive/internal/domain/repo"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	userRepo := repo.NewUserRepo(database.DB)
	userHandler := handlers.NewUserHandler(userRepo, cfg)

	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	r.Static("static", "./static")

	// Public routes
	r.GET("/", handlers.IndexPageHandler)
	r.GET("/login", handlers.LoginPageHandler)
	r.GET("/signup", handlers.SignUpPageHandler)
	r.POST("/signup", userHandler.SingnUpHandler)
	r.POST("/login", userHandler.LoginHandler)

	// Protected routes
	// Protected routes
	protected := r.Group("/")
	protected.Use(func(c *gin.Context) {
		// Проверяем токен в URL только для GET запросов
		if c.Request.Method == "GET" {
			if token := c.Query("token"); token != "" {
				c.Request.Header.Set("Authorization", "Bearer "+token)
			}
		}
		c.Next()
	})
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.GET("/main", handlers.MainPage)
		protected.GET("/profile", handlers.ProfilePageHandler)
	}

	r.Run(":8080")
}
