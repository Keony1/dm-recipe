package protocols

type GifRepository interface {
	Find(title string) (string, error)
}
