package utils

import (
	"github.com/gin-gonic/gin"

	"video/pkg/models"
)

func GetUser(c *gin.Context) *models.User {
	return c.Value("user").(*models.User)
}
