package main

import (
	"bufio"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// Reads a recipe from a file and writes out the corresponding html file using the given template
func RecipeToHtml(recipePath, htmlPath string, htmlTemplate *template.Template) error {
	// Open and read contents
	file, err := os.Open(recipePath)
	if err != nil {
		return err
	}
	defer file.Close()
	recipe := MarshalRecipe(bufio.NewReader(file))

	// Format and write the page to a file
	outFile, err := os.Create(htmlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = htmlTemplate.Execute(outFile, recipe)

	return err
}

// Reads all *.recipe files in recipeDir (except template.recipe if it exists, and generates *.html for each of them in the output directory
func GenerateAllRecipes(recipeDir, outDir string) error {
	// Make output directory if necessary
	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	// Open html template
	htmlTemplate, err := template.ParseFiles("data/templates/recipe.html")
	if err != nil {
		return err
	}

	// Walk recipeDir
	return filepath.Walk(recipeDir, func(recipePath string, f os.FileInfo, err error) error {
		// Make sure recipePath is a normal recipe file, and not template.recipe
		if !f.Mode().IsRegular() {
			return nil
		}
		if filepath.Ext(recipePath) != ".recipe" {
			return nil
		}
		if filepath.Base(recipePath) == "template.recipe" {
			return nil
		}

		htmlFileName := strings.Replace(filepath.Base(recipePath), "recipe", "html", 1)

		return RecipeToHtml(recipePath, filepath.Join(outDir, htmlFileName), htmlTemplate)
	})
}
