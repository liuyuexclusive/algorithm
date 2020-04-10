package bitsum

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{5, 6}, want: 11},
		{args: args{0, 0}, want: 0},
		{args: args{-1, 6}, want: 5},
		{args: args{0, 100}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
