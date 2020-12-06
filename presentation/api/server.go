package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	domain "github.com/keony1/dm-recipe/domain/usecases"
	"github.com/keony1/dm-recipe/presentation/presenter"
)

// ErrKeyWordsLimit is returned when i param is more than 3
var ErrKeyWordsLimit = errors.New("Only 3 ingredients allowed")

// Server is used to handle http
type Server struct {
	http.Handler
	LoadRecipes domain.LoadRecipes
}

// NewServer returns a Server to handle http
func NewServer(lr domain.LoadRecipes) *Server {
	s := new(Server)
	s.LoadRecipes = lr

	r := http.NewServeMux()
	r.HandleFunc("/recipes", s.recipes)

	s.Handler = r

	return s
}

func (s *Server) recipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := r.URL.Query()["i"]

	keywords := []string{}
	var ingredients string
	if len(params) > 0 {
		ingredients = params[0]
		keywords = strings.Split(ingredients, ",")
	}

	if err := checkKeywords(keywords, w); err != nil {
		return
	}

	results, _ := s.LoadRecipes.Load(ingredients)

	var recipes []presenter.Recipe
	for _, result := range results {
		recipes = append(recipes, presenter.Recipe{
			Title:       result.Title,
			Ingredients: result.Ingredients,
			Link:        result.Link,
			Gif:         result.Gif,
		})
	}

	response := presenter.Response{
		Keywords: keywords,
		Recipes:  recipes,
	}

	m, _ := json.Marshal(response)

	fmt.Fprint(w, string(m))
}

func checkKeywords(keywords []string, w http.ResponseWriter) error {
	if len(keywords) > 3 {
		w.WriteHeader(http.StatusBadRequest)

		kwError := presenter.Error{
			Message: ErrKeyWordsLimit.Error(),
		}
		errJSON, _ := json.Marshal(kwError)

		fmt.Fprint(w, string(errJSON))

		return ErrKeyWordsLimit
	}

	return nil
}
