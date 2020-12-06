package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "github.com/keony1/dm-recipe/domain/usecases"
	"github.com/keony1/dm-recipe/presentation/presenter"
)

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

	var ingredients string
	if len(params) > 0 {
		ingredients = params[0]
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

	json.NewEncoder(w).Encode(recipes)
	fmt.Fprint(w, recipes)
}
