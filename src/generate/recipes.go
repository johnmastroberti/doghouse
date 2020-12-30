package generate

import (
	"html/template"
	"os"
	"path/filepath"

	"doghousecooking.com/sitegen/src/recipe"
)

// Reads a recipe from a file and writes out the corresponding html file using the given template
func recipeToHtml(r recipe.Recipe, htmlPath string, htmlTemplate *template.Template) error {
	// Format and write the page to a file
	outFile, err := os.Create(htmlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = htmlTemplate.Execute(outFile, r)

	return err
}

// Reads all *.recipe files in recipeDir (except template.recipe if it exists, and generates *.html for each of them in the output directory
func AllRecipes(book recipe.RecipeBook, outDir string) error {
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

	// Generate the webpages
	for _, r := range book {
		err = recipeToHtml(r, filepath.Join(outDir, r.PageName()), htmlTemplate)
		if err != nil {
			return err
		}
	}
	return nil
}
