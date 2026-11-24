package main

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	want := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	if got := FloodFill(img, 1, 1, 2); !reflect.DeepEqual(got, want) {
		t.Errorf("FloodFill = %v, want %v", got, want)
	}

	// Same color: untouched, and must not loop forever.
	same := [][]int{{0, 0}, {0, 0}}
	if got := FloodFill(same, 0, 0, 0); !reflect.DeepEqual(got, [][]int{{0, 0}, {0, 0}}) {
		t.Errorf("same color fill changed image: %v", got)
	}

	// Diagonals don't connect.
	diag := [][]int{{1, 0}, {0, 1}}
	want2 := [][]int{{9, 0}, {0, 1}}
	if got := FloodFill(diag, 0, 0, 9); !reflect.DeepEqual(got, want2) {
		t.Errorf("diagonal leak: %v", got)
	}

	if got := FloodFill([][]int{}, 0, 0, 5); len(got) != 0 {
		t.Error("empty image")
	}
}
