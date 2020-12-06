package main

import (
	"net/http"

	data "github.com/keony1/dm-recipe/data/usecases"
	"github.com/keony1/dm-recipe/infra/remote"
	"github.com/keony1/dm-recipe/presentation/api"
)

func main() {
	gr := remote.GifRepository{}
	rr := remote.RecipesRepository{}

	remote := data.NewRemoteLoadRecipes(gr, rr)

	s := api.NewServer(remote)
	http.ListenAndServe(":8080", s)
}
