package main

import (
	"log"

	"doghousecooking.com/sitegen/src/generate"
)

func main() {
	// Generate recipes
	check(generate.AllRecipes("data/recipes", "www/recipes"))
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
