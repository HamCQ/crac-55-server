package entities

import (
	"reflect"
	"testing"
)

func TestSearchBxCraMap(t *testing.T) {
	tests := []struct {
		name string
		want []SearchBxCraInfo
	}{
		// TODO: Add test cases.
		{
			name: "t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchBxCraPointMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchBxCraMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
