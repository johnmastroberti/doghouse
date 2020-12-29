package main

import (
	"log"
)

func main() {
	// Generate recipes
	check(GenerateAllRecipes("data/recipes", "www/recipes"))
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
