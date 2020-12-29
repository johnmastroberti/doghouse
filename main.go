package main

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Recipe struct {
	Title       string
	Description string
	PrepTime    string
	CookTime    string
	Yield       string
	Ingredients []template.HTML
	Directions  []template.HTML
	Notes       []template.HTML
}

func MarshalRecipe(recipe *bufio.Reader) Recipe {
	var out Recipe
	inIngredients := false
	inDirections := false
	inNotes := false
	for {
		line, err := recipe.ReadString('\n')
		if err != nil {
			return out
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch strings.Fields(line)[0] {
		case "TITLE:":
			out.Title = line[7:]
			continue
		case "DESCRIPTION:":
			out.Description = line[13:]
			continue
		case "PREP:":
			out.PrepTime = line[6:]
			continue
		case "COOK:":
			out.CookTime = line[6:]
			continue
		case "YIELD:":
			out.Yield = line[7:]
			continue
		case "INGREDIENTS:":
			inIngredients = true
			continue
		case "DIRECTIONS:":
			inIngredients = false
			inDirections = true
			continue
		case "NOTES:":
			inDirections = false
			inNotes = true
			continue
		}

		switch {
		case inIngredients:
			// Parse "-- blah" as <h3>blah</h3>
			if strings.Index(line, "-- ") == 0 && len(line) > 3 {
				out.Ingredients = append(out.Ingredients, template.HTML("<h3>"+line[3:]+"</h3>"))
			} else {
				out.Ingredients = append(out.Ingredients, template.HTML(line))
			}
		case inDirections:
			out.Directions = append(out.Directions, template.HTML(line))
		case inNotes:
			out.Notes = append(out.Notes, template.HTML(line))
		}

	}

	return out
}

func main() {
	// Load the template
	htmlTemplate, err := template.ParseFiles("template.html")
	check(err)

	// Make the html directory if necessary
	err = os.MkdirAll("html", 0755)
	check(err)

	// Walk the recipes directory
	recipeDir := "raw"
	err = filepath.Walk(recipeDir, func(path string, f os.FileInfo, err error) error {
		// Make sure it is a normal recipe file
		if !f.Mode().IsRegular() {
			return nil
		}
		if filepath.Ext(path) != ".recipe" {
			return nil
		}

		// Open and read contents
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		recipe := MarshalRecipe(bufio.NewReader(file))

		// Format and write the page to a file
		outName := filepath.Join("html", strings.Replace(filepath.Base(path), "recipe", "html", 1))
		outFile, err := os.Create(outName)
		if err != nil {
			return err
		}
		defer outFile.Close()
		err = htmlTemplate.Execute(outFile, recipe)

		return err
	})
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
