package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/user-register", h.createUser)
		auth.POST("/user-login", h.loginUser)
	}

	apiV1 := router.Group("/api/v1", h.userIdentity)
	{
		users := apiV1.Group("/users")
		{
			users.PUT("/", h.updateUser)
			users.DELETE("/", h.deleteUser)
		}

	}

	return router
}
