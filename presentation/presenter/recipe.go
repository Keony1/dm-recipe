package presenter

// Recipe data
type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

type Response struct {
	Keywords []string `json:"keywords"`
	Recipes  []Recipe `json:"recipes"`
}
