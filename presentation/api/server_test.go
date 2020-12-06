package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/keony1/dm-recipe/domain/entities"
	"github.com/keony1/dm-recipe/presentation/presenter"
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

		var r presenter.Response
		json.Unmarshal(res.Body.Bytes(), &r)

		assertKeyWords(t, r.Keywords, nil)
		assertContentType(t, res, jsonContentType)
		assertStatusCode(t, res.Code, http.StatusOK)
	})
}

func assertKeyWords(t *testing.T, got, want []string) {
	t.Helper()

	if reflect.DeepEqual(got, want) {
		t.Errorf("expecting response.Keywords to be %v, but got %v", got, want)
	}
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
