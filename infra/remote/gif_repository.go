package remote

import (
	"fmt"
	"net/http"
	"net/url"
)

const giphyURL = "http://api.giphy.com/v1/gifs/search"

// GifRepository is the struct used to fetch data from giphy api
type GifRepository struct{}

// Find is the implementation of protocol defined in data layer of GifRepository
func (g GifRepository) Find(title string) (string, error) {
	qt := url.QueryEscape(title)

	resp, err := http.Get(fmt.Sprintf("%v?q=%v&api_key=s9GY9cDxXGHLCVCZifd0lh0rP8x1o2lh", giphyURL, qt))

	if err != nil {
		return "", err
	}

	var gr gifResponse
	convertToJSON(&gr, resp)

	if len(gr.Data) > 0 {
		gif := gr.Data[0]
		return gif.URL, nil
	}

	return "", nil
}

type gifResponse struct {
	Data []gif
}

type gif struct {
	URL string
}
