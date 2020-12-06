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

var ErrKeyWordsLimit = errors.New("Only 3 ingredients allowed")

type Server struct {
	http.Handler
	LoadRecipes domain.LoadRecipes
}

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

	if len(keywords) > 3 {
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(ErrKeyWordsLimit)

		fmt.Fprint(w, m)
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
