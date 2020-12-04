package protocols

type PuppyRepository interface {
	Load(search string) ([]PuppyResult, error)
}

type PuppyResult struct {
	Title, Href, Ingredients string
}
