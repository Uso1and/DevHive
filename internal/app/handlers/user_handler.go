package handlers

import (
	"database/sql"
	"devhive/internal/domain/config"
	"devhive/internal/domain/models"
	"devhive/internal/domain/repo"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRep repo.UserRepoInterface
	cfg     *config.ConfigDB
}

func NewUserHandler(userRep repo.UserRepoInterface, cfg *config.ConfigDB) *UserHandler {
	return &UserHandler{userRep: userRep, cfg: cfg}
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

	user.Password = ""
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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Printf("LoginHandler - password mismatch: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		log.Printf("LoginHandler - error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	log.Printf("Generated token for user %s: %s", user.Username, tokenString)

	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"expires": expirationTime,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
