package usecases

import (
	"reflect"
	"testing"
)

func TestRemoteLoadRecipes(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rl := NewRemoteLoadRecipes(tt.spyGif, tt.spyPuppy)
			got, err := rl.Load(tt.args)

			if tt.expectError {
				// expecting an error
				if err == nil {
					t.Fatalf("RemoteLoadRecipes.Load(%q); expecting error but got nil", tt.args)
				}
			} else {
				// not expecting an error
				if err != nil {
					t.Fatalf("RemoteLoadRecipes.Load(%q); return unexpected error: %v", tt.args, err)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RemoteLoadRecipes.Load(%q) = %v; got %v", tt.args, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkRemoteLoadRecipes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			rl := NewRemoteLoadRecipes(tt.spyGif, tt.spyPuppy)
			rl.Load(tt.args)
		}
	}
}
