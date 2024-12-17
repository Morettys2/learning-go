package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GET Ok"})
}

func Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "SHOW Ok"})
}

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Ok"})
}

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Upatede Ok"})
}

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List Ok"})
}
