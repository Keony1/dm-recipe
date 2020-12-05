package protocols

// GifRepository returns respective url for a given title
type GifRepository interface {
	Find(title string) (string, error)
}
