package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

	var ppResult []protocols.PuppyResult

	for _, recipe := range ppResponse.Results {
		xsIngredients := strings.Split(recipe.Ingredients, ",")

		pr := protocols.PuppyResult{
			Title:       recipe.Title,
			Href:        recipe.Href,
			Ingredients: xsIngredients,
		}

		ppResult = append(ppResult, pr)
	}

	return ppResult, nil
}
