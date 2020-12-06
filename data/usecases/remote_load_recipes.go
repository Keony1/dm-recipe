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
	ppRecipes, err := r.puppyRepository.Load(search)

	if err != nil {
		return nil, err
	}

	var recipes []*entities.Recipe
	for _, ppRecipe := range ppRecipes {
		gifURL, err := r.gifRepository.Find(ppRecipe.Title)

		if err != nil {
			return nil, err
		}

		recipe := &entities.Recipe{
			Title:       ppRecipe.Title,
			Ingredients: ppRecipe.Ingredients,
			Link:        ppRecipe.Href,
			Gif:         gifURL,
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
