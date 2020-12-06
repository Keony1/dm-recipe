package remote

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/keony1/dm-recipe/config"
)

const giphyURL = "https://api.giphy.com/v1/gifs/search"

// GifRepository is the struct used to fetch data from giphy api
type GifRepository struct{}

// Find is the implementation of protocol defined in data layer of GifRepository
func (g GifRepository) Find(title string) (string, error) {
	qt := url.QueryEscape(title)

	resp, err := http.Get(fmt.Sprintf("%v?api_key=%v&q=%v&", giphyURL, config.GIPHY_KEY, qt))

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
