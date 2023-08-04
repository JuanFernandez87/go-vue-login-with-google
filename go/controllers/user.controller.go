package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JuanFernandez87/go-vue-login-with-google/models"
)

func GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(&currentUser)}})
}
