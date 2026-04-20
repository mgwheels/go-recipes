package routes

import (
	"go-recipes/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/recipes", handlers.GetAllRecipes)
		api.GET("/recipes/:id", handlers.GetRecipe)
		api.POST("/recipes", handlers.CreateRecipe)
		api.PUT("/recipes/:id", handlers.UpdateRecipe)
		api.DELETE("/recipes/:id", handlers.DeleteRecipe)
	}
}
