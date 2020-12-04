package usecases

import (
	"reflect"
	"testing"

	"github.com/keony1/dm-recipe/data/protocols"
	"github.com/keony1/dm-recipe/domain/entities"
)

func TestRemoteLoadRecipes(t *testing.T) {
	resultPp := []protocols.PuppyResult{{Title: "any_title"}}
	spyGif := SpyGifRepository{result: "any_gif", err: nil}
	spyPp := SpyPuppyRepository{result: resultPp, err: nil}

	rl := NewRemoteLoadRecipes(spyGif, spyPp)
	recipes, _ := rl.Load("any_search")
	want := []*entities.Recipe{{Title: "any_title", Gif: "any_gif"}}

	if !reflect.DeepEqual(recipes, want) {
		t.Errorf("recipes.Load() = %v, but want %v", recipes, want)
	}
}

type SpyGifRepository struct {
	result string
	err    error
}

func (s SpyGifRepository) Find(title string) (string, error) {
	return s.result, s.err
}

type SpyPuppyRepository struct {
	result []protocols.PuppyResult
	err    error
}

func (p SpyPuppyRepository) Load(search string) ([]protocols.PuppyResult, error) {
	return p.result, p.err
}
