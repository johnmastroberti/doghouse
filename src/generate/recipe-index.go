package generate

import (
	"html/template"
	"os"
	"path/filepath"

	"doghousecooking.com/sitegen/src/recipe"
)

// Generates index.html for the recipe directory
func RecipeIndex(book recipe.RecipeBook, outDir string) error {
	// Open the template
	htmlTemplate, err := template.ParseFiles("data/templates/recipe-index.html")
	if err != nil {
		return err
	}

	// Prepare the index file for writing
	outFile, err := os.Create(filepath.Join(outDir, "index.html"))
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = htmlTemplate.Execute(outFile, book)

	return err
}
