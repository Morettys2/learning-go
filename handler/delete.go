package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Ok"})
}
