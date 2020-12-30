package main

import (
	"log"

	"doghousecooking.com/sitegen/src/generate"
	"doghousecooking.com/sitegen/src/recipe"
)

func main() {
	// Parse recipes
	book := recipe.NewBook("data/recipes")

	// Generate recipes
	check(generate.AllRecipes(book, "www/recipes"))
	check(generate.RecipeIndex(book, "www/recipes"))
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
