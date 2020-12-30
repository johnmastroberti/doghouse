package recipe

import (
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Recipe represents all of the information about a single recipe which is encoded in a recipe file.
type Recipe struct {
	Title       string
	Description string
	PrepTime    string
	CookTime    string
	Yield       string
	Image       string
	Ingredients []template.HTML // These are stored as HTML to allow inline
	Directions  []template.HTML // html in the recipe files
	Notes       []template.HTML
}

// Returns the name of the webpage for the given recipe
func (r *Recipe) PageName() string {
	// Use Title, lowercased with spaces set to dashes
	name := strings.ToLower(r.Title)
	name = strings.ReplaceAll(name, " ", "-")

	reg := regexp.MustCompile("[^a-z-]+")
	name = reg.ReplaceAllString(name, "")

	return name + ".html"
}

// A RecipeBook holds all of the recipes in memory, to be passed around instead of being read in from disk multiple times
type RecipeBook []Recipe

// Reads all files *.recipe (except template.recipe) and returns a RecipeBook containing all of them
func NewBook(recipeDir string) RecipeBook {
	var book RecipeBook

	filepath.Walk(recipeDir, func(recipePath string, f os.FileInfo, err error) error {
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

		r, err := ParseFile(recipePath)
		if err == nil {
			book = append(book, r)
		}
		return err
	})

	return book
}
