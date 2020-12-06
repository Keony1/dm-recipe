package usecases

import "github.com/keony1/dm-recipe/domain/entities"

// LoadRecipes interface
type LoadRecipes interface {
	Load(string) ([]*entities.Recipe, error)
}
