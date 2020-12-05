package protocols

// PuppyRepository loads recipes from puppy api
type PuppyRepository interface {
	Load(search string) ([]PuppyResult, error)
}

// PuppyResult is the result from puppy api
type PuppyResult struct {
	Title       string
	Href        string
	Ingredients []string
}
