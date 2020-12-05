package remote

import (
	"reflect"
	"testing"
)

func TestRecipesRepository(t *testing.T) {
	r := RecipesRepository{}

	t.Run("empty response", func(t *testing.T) {
		got, err := r.Load("1")

		checkError(t, err, nil)
		checkLen(t, len(got), 0)
	})

	t.Run("verify ingredients", func(t *testing.T) {
		resp, err := r.Load("banana")
		want := []string{"banana"}

		checkError(t, err, nil)

		got := resp[0].Ingredients
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expect an array of string %v, but got %v", want, got)
		}
	})
}

func TestGifRepository(t *testing.T) {
	r := GifRepository{}

	t.Run("empty response", func(t *testing.T) {
		got, err := r.Find("")
		want := ""

		checkError(t, err, nil)
		checkResponse(t, got, want)
	})
}

func checkError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Fatalf("RecipesRepository.Load() error = %v, wantErr %v", got, want)
	}
}

func checkLen(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("RecipesRepository.Load() have len %d; but want %d", got, want)
	}
}

func checkResponse(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q; but want %q", got, want)
	}
}
