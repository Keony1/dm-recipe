package usecases

import (
	"errors"

	"github.com/keony1/dm-recipe/data/protocols"
	"github.com/keony1/dm-recipe/domain/entities"
)

var (
	ErrGifRepository   = errors.New("Somethig is wrong with Giphy api")
	ErrPuppyRepository = errors.New("Somethig is wrong with Puppy api")
)

var testCases = []struct {
	name        string
	args        string
	spyPuppy    SpyPuppyRepository
	spyGif      SpyGifRepository
	want        []*entities.Recipe
	expectError bool
}{
	{
		name: "loaded correctly",
		args: "any_search",
		spyPuppy: SpyPuppyRepository{
			result: []protocols.PuppyResult{
				{
					Title: "any_title",
					Href:  "any_url",
				},
			},
			err: nil,
		},
		spyGif: SpyGifRepository{
			result: "any_url_gif",
			err:    nil,
		},
		want: []*entities.Recipe{
			{
				Title: "any_title",
				Link:  "any_url",
				Gif:   "any_url_gif",
			},
		},
		expectError: false,
	},
	{
		name: "returns error if PuppyRepository returns error",
		args: "",
		spyPuppy: SpyPuppyRepository{
			result: nil,
			err:    ErrPuppyRepository,
		},
		spyGif:      SpyGifRepository{"", nil},
		want:        nil,
		expectError: true,
	},
	{
		name: "returns error if GifRepository returns error",
		args: "incorrect_title",
		spyPuppy: SpyPuppyRepository{
			result: []protocols.PuppyResult{
				{
					Title: "any_title",
					Href:  "any_url",
				},
			},
			err: nil,
		},
		spyGif: SpyGifRepository{
			result: "",
			err:    ErrGifRepository,
		},
		want:        nil,
		expectError: true,
	},
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
