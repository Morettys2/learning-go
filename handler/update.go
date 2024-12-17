package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Upatede Ok"})
}
