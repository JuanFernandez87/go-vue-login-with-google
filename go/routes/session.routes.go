package routes

import (
	"github.com/JuanFernandez87/go-vue-login-with-google/controllers"
	"github.com/gin-gonic/gin"
)

type SessionRouteController struct {
	authController controllers.AuthController
}

func NewSessionRouteController(authController controllers.AuthController) SessionRouteController {
	return SessionRouteController{authController}
}

func (rc *SessionRouteController) SessionRoute(rg *gin.RouterGroup) {
	router := rg.Group("/sessions/oauth")

	router.GET("/google", rc.authController.GoogleOAuth)
}
