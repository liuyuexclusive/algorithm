package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{[]int{1, 2, 3, 1}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
