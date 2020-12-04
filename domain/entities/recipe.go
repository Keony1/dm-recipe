package entities

// Recipe data
type Recipe struct {
	Title       string
	Ingredients []string
	Link        string
	Gif         string
}

// NewRecipe returns a pointer to a Recipe
func NewRecipe(title, link, gif string, ingredients []string) *Recipe {
	r := &Recipe{
		Title:       title,
		Link:        link,
		Gif:         link,
		Ingredients: ingredients,
	}

	return r
}
