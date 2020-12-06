package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keony1/dm-recipe/domain/entities"
)

func TestServer_recipes(t *testing.T) {
	t.Run("without query param", func(t *testing.T) {
		loadResults := []*entities.Recipe{
			{
				Title: "ice cream",
			},
		}

		spyLoadRecipes := SpyLoadRecipes{results: loadResults, err: nil}

		server := NewServer(spyLoadRecipes)

		req, _ := http.NewRequest(http.MethodGet, "/recipes", nil)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)

		assertContentType(t, res, jsonContentType)
		assertStatusCode(t, res.Code, http.StatusOK)
	})
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("expect status code to be %d, but got %d", want, got)
	}
}

const jsonContentType = "application/json"

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

type SpyLoadRecipes struct {
	results []*entities.Recipe
	err     error
}

func (s SpyLoadRecipes) Load(title string) ([]*entities.Recipe, error) {
	return s.results, s.err
}
