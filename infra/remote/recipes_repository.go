package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/keony1/dm-recipe/data/protocols"
)

type RecipesRepository struct{}

func (r RecipesRepository) Load(search string) ([]protocols.PuppyResult, error) {

	resp, err := http.Get(fmt.Sprintf("http://www.recipepuppy.com/api/?i=%v", search))

	if err != nil {
		return nil, err
	}

	dataBytes, _ := ioutil.ReadAll(resp.Body)

	var ppResponse protocols.PuppyResponse

	jsonConverterErr := json.Unmarshal([]byte(dataBytes), &ppResponse)

	if jsonConverterErr != nil {
		return nil, jsonConverterErr
	}

	return nil, nil
}
