package remote

import (
	"reflect"
	"testing"

	"github.com/keony1/dm-recipe/data/protocols"
)

func TestRecipesRepository(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name    string
		args    args
		want    []protocols.PuppyResult
		wantErr bool
	}{
		{
			name:    "empty PuppyResponse",
			args:    args{"1"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RecipesRepository{}
			got, err := r.Load(tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecipesRepository.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecipesRepository.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
