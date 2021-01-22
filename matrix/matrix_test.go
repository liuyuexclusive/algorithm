package matrix

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{args: args{[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			want := [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			}
			if !reflect.DeepEqual(tt.args.matrix, want) {
				t.Errorf("want: %v, got: %v", want, tt.args.matrix)
			}
		})
	}
}

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{args: args{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			want := [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			}
			if !reflect.DeepEqual(tt.args.matrix, want) {
				t.Errorf("want: %v,got: %v", want, tt.args.matrix)
			}
		})
	}
}

func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args: args{board: board, word: "ABCCED"}, want: true},
		// {args: args{board: board, word: "SEE"}, want: true},
		// {args: args{board: board, word: "ABCB"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
