package handlers

import (
	"database/sql"
	"devhive/internal/domain/models"
	"devhive/internal/domain/repo"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRep repo.UserRepoInterface
}

func NewUserHandler(userRep repo.UserRepoInterface) *UserHandler {
	return &UserHandler{userRep: userRep}
}

func (h *UserHandler) SingnUpHandler(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if request.Username == "" || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username, email and password are required"})
		return
	}

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	user := models.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hashpassword),
		CreatedAt: time.Now(),
	}

	if err := h.userRep.CreateUser(c.Request.Context(), &user); err != nil {
		log.Printf("Error create user:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	user.Password = "" // Не возвращаем пароль в ответе
	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"user":    user,
	})
}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Printf("LoginHandler - invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	log.Printf("Attempting login for user: %s", credentials.Username)

	user, err := h.userRep.GetUserByUsername(c.Request.Context(), credentials.Username)
	if err != nil {
		log.Printf("LoginHandler - error getting user: %v", err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	log.Printf("User found: %+v", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Printf("LoginHandler - password mismatch: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	log.Printf("Login successful for user: %s", user.Username)

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"user":    user,
	})
}
