package main

import (
	"reflect"
	"testing"
)

func Test_bubblesort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{[]int{1, 4, 3, 2}}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubblesort(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubblesort() = %v, want %v", got, tt.want)
			}
		})
	}
}
