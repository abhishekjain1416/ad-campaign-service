package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}
