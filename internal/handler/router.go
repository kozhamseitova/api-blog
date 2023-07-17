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
			users.GET("/:user_id/articles", h.getArticleByUserId)
		}

		articles := apiV1.Group("/articles")
		{
			articles.POST("/", h.createArticle)
			articles.GET("/", h.getAllArticles)
			articles.PUT("/:id", h.updateArticle)
			articles.DELETE("/:id", h.deleteArticle)
			articles.GET("/:id", h.getArticleById)
		}

		categories := apiV1.Group("/categories")
		{
			categories.GET("/", h.getAllCategories)
		}

	}

	return router
}
