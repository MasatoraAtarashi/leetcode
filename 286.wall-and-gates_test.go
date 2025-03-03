package main

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  [][]int
	}{
		{
			name: "multiple rooms with walls and gates",
			rooms: [][]int{
				{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "single room with wall",
			rooms: [][]int{
				{-1},
			},
			want: [][]int{
				{-1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([][]int, len(tt.rooms))
			for i := range tt.rooms {
				original[i] = make([]int, len(tt.rooms[i]))
				copy(original[i], tt.rooms[i])
			}

			wallsAndGates(tt.rooms)

			if !reflect.DeepEqual(tt.rooms, tt.want) {
				t.Errorf("wallsAndGates() = %v, want %v", tt.rooms, tt.want)
			}
		})
	}
}
