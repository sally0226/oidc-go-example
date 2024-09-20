package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(h *handler) *gin.Engine {
	router := gin.Default()

	// API Docs
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/login/:provider", h.OAuthLogin)
		authGroup.GET("/callback", h.OAuthCallback)
	}

	return router
}
