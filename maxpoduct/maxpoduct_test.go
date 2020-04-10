package maxpoduct

import "testing"

func TestMaxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 3, -2, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct(tt.args.nums); got != tt.want {
				t.Errorf("MaxProduct2() = %v, want %v", got, tt.want)
			}
		})
	}
}
