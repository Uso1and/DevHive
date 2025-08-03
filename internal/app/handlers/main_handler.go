package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func SignUpPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"userID":   c.MustGet("userID"),
		"username": c.MustGet("username"),
		"message":  "Данные",
	})

}

func ProfilePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"userID":   c.MustGet("userID"),
		"username": c.MustGet("username"),
		"message":  "профиль",
	})
}
