package remote

import (
	"fmt"
	"net/http"
	"net/url"
)

type GifRepository struct{}

func (g GifRepository) Find(title string) (string, error) {
	qt := url.QueryEscape(title)

	resp, err := http.Get(fmt.Sprintf("http://api.giphy.com/v1/gifs/search?q=%v&api_key=s9GY9cDxXGHLCVCZifd0lh0rP8x1o2lh", qt))

	if err != nil {
		return "", err
	}

	var gr gifResponse
	convertToJSON(gr, resp)

	if gr.data != nil {
		gif := gr.data[0]
		return gif.url, nil
	}

	return "", nil
}

type gifResponse struct {
	data []gif
}

type gif struct {
	url string
}
