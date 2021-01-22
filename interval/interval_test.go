package interval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{args: args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, want: [][]int{{1, 5}, {6, 9}}},
		{args: args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, want: [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{args: args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{args: args{[][]int{{1, 4}, {4, 5}}}, want: [][]int{{1, 5}}},
		{args: args{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}}, want: [][]int{{1, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
