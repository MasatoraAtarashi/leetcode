package main

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "Example 1",
		// 	args: args{
		// 		board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		// 		word:  "ABCCED",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{
		// 		board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		// 		word:  "SEE",
		// 	},
		// 	want: true,
		// },
		{
			name: "Example 3",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := exist(tt.args.board, tt.args.word); got != tt.want {
			t.Errorf("%q. exist() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
