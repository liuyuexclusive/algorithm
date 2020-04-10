package maxprofit

import (
	"testing"
)

func TestMaxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProfit1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit1(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit1() = %v, want %v", got, tt.want)
			}
		})
	}
}
