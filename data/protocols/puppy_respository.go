package protocols

// PuppyRepository loads recipes from puppy api
type PuppyRepository interface {
	Load(search string) ([]PuppyResult, error)
}

// PuppyResponse is the slice of results returned from puppy api
type PuppyResponse struct {
	Results []PuppyRecipe
}

//PuppyRecipe is the actual recipe returned from puppy api
type PuppyRecipe struct {
	Title, Href, Ingredients string
}

// PuppyResult is the parsed []PuppyResponse
type PuppyResult struct {
	Title       string
	Href        string
	Ingredients []string
}
