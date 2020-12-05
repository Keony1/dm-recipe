package usecases

import (
	"github.com/keony1/dm-recipe/data/protocols"
	"github.com/keony1/dm-recipe/domain/entities"
)

// RemoteLoadRecipes struct
type RemoteLoadRecipes struct {
	gifRepository   protocols.GifRepository
	puppyRepository protocols.PuppyRepository
}

// Load remotely
func (r *RemoteLoadRecipes) Load(search string) ([]*entities.Recipe, error) {
	ppRecipes, _ := r.puppyRepository.Load(search)

	var recipes []*entities.Recipe
	for _, ppRecipe := range ppRecipes {
		gifURL, _ := r.gifRepository.Find(ppRecipe.Title)

		recipe := &entities.Recipe{
			Title: ppRecipe.Title,
			Link:  ppRecipe.Href,
			Gif:   gifURL,
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// NewRemoteLoadRecipes returns new RemoteLoadRecipes
func NewRemoteLoadRecipes(g protocols.GifRepository, p protocols.PuppyRepository) *RemoteLoadRecipes {
	r := &RemoteLoadRecipes{
		gifRepository:   g,
		puppyRepository: p,
	}

	return r
}
