package recipe

import (
	"bufio"
	"html/template"
	"os"
	"strings"
)

// Reads a recipe file and returns
func MarshalRecipe(recipe *bufio.Reader) (Recipe, error) {
	var out Recipe
	out.Image = "placeholder.jpg"
	inIngredients := false
	inDirections := false
	inNotes := false
	for {
		// Read a line
		line, err := recipe.ReadString('\n')
		if err != nil {
			break
			// if err == io.EOF {
			//   err = nil
			// }
			// return out, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Handle keywords
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
		case "IMAGE:":
			out.Image = line[7:]
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

		// Add lines to the appropriate section
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

	// Finally, fill the PageName field
	out.PageName = out.makePageName()

	return out, nil
}

// Reads a recipe file and returns the corresponding recipe
func ParseFile(recipePath string) (Recipe, error) {
	// Open and read contents
	file, err := os.Open(recipePath)
	if err != nil {
		return Recipe{}, err
	}
	defer file.Close()
	return MarshalRecipe(bufio.NewReader(file))
}
