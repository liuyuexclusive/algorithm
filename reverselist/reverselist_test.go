package reverselist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head := &ListNode{Val: 1}
	head.Add(2).Add(3).Add(4).Add(5)
	head2 := &ListNode{Val: 5}
	head2.Add(4).Add(3).Add(2).Add(1)
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{head}, want: head2.show()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got.show(), tt.want) {
				t.Errorf("reverseList() = %v, want %v", got.show(), tt.want)
			}
		})
	}
}
