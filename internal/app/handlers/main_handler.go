package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)

}

func SignUpPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func ProfilePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", nil)
}
