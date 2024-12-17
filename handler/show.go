package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "SHOW Ok"})
}
