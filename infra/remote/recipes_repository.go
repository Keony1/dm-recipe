package remote

import (
	"fmt"
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

	var ppResponse protocols.PuppyResponse
	convertToJSON(&ppResponse, resp)
	parsedRecipes := parseResults(ppResponse.Results)

	return parsedRecipes, nil
}

func parseResults(recipes []protocols.PuppyRecipe) []protocols.PuppyResult {
	var ppResult []protocols.PuppyResult
	for _, recipe := range recipes {
		xsIngredients := strings.Split(recipe.Ingredients, ",")

		pr := protocols.PuppyResult{
			Title:       recipe.Title,
			Href:        recipe.Href,
			Ingredients: xsIngredients,
		}

		ppResult = append(ppResult, pr)
	}

	return ppResult
}
