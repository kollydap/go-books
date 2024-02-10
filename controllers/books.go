//controllers/ books.go

package controllers

import (
	"net/http"

	"core/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
