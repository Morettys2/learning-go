package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List Ok"})
}
