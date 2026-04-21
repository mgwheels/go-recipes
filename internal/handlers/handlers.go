package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-recipes/internal/models"
)

// Local recipe store
var recipes []models.Recipe

func GetAllRecipes(c *gin.Context) {
	c.JSON(200, recipes)
}

func GetRecipe(c *gin.Context) {
	// Get ID from parameter
	id := c.Param("id")

	// Loop through recipe store
	for _, r := range recipes {
		if r.Id == id {
			c.JSON(200, r)
			return
		}
	}

	// Recipe not found in recipe store, return err
	c.JSON(404, gin.H{"error": fmt.Sprintf("recipe not found for id: %s", id)})
}

func CreateRecipe(c *gin.Context) {
	var newRecipe models.Recipe

	// c.ShouldBindJSON is equal to something like json.Unmarshal(body, &newRecipe).
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Generate new ID for recipe
	recipeId, err := uuid.NewV7()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("error generating uuid: %s", err)})
	}
	newRecipe.Id = recipeId.String()

	recipes = append(recipes, newRecipe)
	c.JSON(201, newRecipe)
}

func UpdateRecipe(c *gin.Context) {
	// Get ID from parameter
	id := c.Param("id")

	// Create updated recipe
	var updatedRecipe models.Recipe
	if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find and update recipe from ID
	for i, r := range recipes {
		if r.Id == id {
			updatedRecipe.Id = id // Preserve id
			recipes[i] = updatedRecipe
			c.JSON(200, updatedRecipe)
			return
		}
	}

	// Else, return 404 recipe not found
	c.JSON(404, gin.H{"error": fmt.Sprintf("recipe not found for id: %s", id)})
}

func DeleteRecipe(c *gin.Context) {
	// Get ID from parameter
	id := c.Param("id")

	// If match on ID, set recipes equal to everything before and after recipe at index i, removing recipe[i]
	for i, r := range recipes {
		if r.Id == id {
			recipes = append(recipes[:i], recipes[i+1:]...)
			c.JSON(200, gin.H{"message": "recipe deleted successfully"})
			return
		}
	}

	// Else, return 404 recipe not found
	c.JSON(404, gin.H{"error": fmt.Sprintf("recipe not found for id: %s", id)})
}
