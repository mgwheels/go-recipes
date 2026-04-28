package utils

import (
	"slices"

	"go-recipes/internal/models"
)

func FilterRecipes(recipes *[]models.Recipe, filteredRecipes *[]models.Recipe, ingredients []string, name string) {
	// If both filters are not set, return all recipes
	if name == "" && len(ingredients) <= 0 {
		// Dereference filteredRecipes and copy recipes
		*filteredRecipes = make([]models.Recipe, len(*recipes))
		copy(*filteredRecipes, *recipes)
		return
	}

	// First filter and return recipes by name
	if name != "" {
		for _, r := range *recipes {
			if r.Name == name {
				*filteredRecipes = append(*filteredRecipes, r)
			}
		}
		return
	}

	// Else filter and return recipes by ingredients
	// Grabs all recipes that contain ANY of ingredients
	if len(ingredients) > 0 {
		for _, r := range *recipes {
			for _, ing := range r.Ingredients {
				if slices.Contains(ingredients, ing.Name) {
					*filteredRecipes = append(*filteredRecipes, r)
					break
				}
			}
		}
		return
	}
}
