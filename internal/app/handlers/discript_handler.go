package handlers

import (
	"devhive/internal/domain/models"
	"devhive/internal/domain/repo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DiscHandler struct {
	discRepo repo.DiscRepoInterface
	userRepo repo.UserRepoInterface
}

func NewDiscHandler(dr repo.DiscRepoInterface, ur repo.UserRepoInterface) *DiscHandler {

	return &DiscHandler{
		discRepo: dr,
		userRepo: ur,
	}

}

func (h *DiscHandler) CreareDisc(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var discussion models.Discussion
	if err := c.ShouldBindJSON(&discussion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	discussion.CreatorID = userID.(int)
	discussion.CreatedAt = time.Now()

	if err := h.discRepo.CreateDisc(c.Request.Context(), &discussion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, discussion)
}
