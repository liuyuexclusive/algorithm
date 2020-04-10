package missingnumber

import (
	"testing"
)

func Test_missingNumber1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber1(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumber1([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}
