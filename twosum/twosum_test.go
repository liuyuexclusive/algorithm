package twosum

import (
	"reflect"
	"testing"
)

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//
func TestTwoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{[]int{2, 7, 11, 15}, 9}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "", args: args{[]int{2, 7, 11, 15}, 9}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum1([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum2([]int{2, 7, 11, 15}, 9)
	}
}
