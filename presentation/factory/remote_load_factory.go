package factory

import (
	"github.com/keony1/dm-recipe/data/usecases"
	"github.com/keony1/dm-recipe/infra/remote"
)

// RemoteLoadRecipesFactory returns an instance of RemoteLoadRecipes
func RemoteLoadRecipesFactory() *usecases.RemoteLoadRecipes {
	gr := remote.GifRepository{}
	rr := remote.RecipesRepository{}
	remote := usecases.NewRemoteLoadRecipes(gr, rr)

	return remote
}
