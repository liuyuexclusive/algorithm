package productexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
