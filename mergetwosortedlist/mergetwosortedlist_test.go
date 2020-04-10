package mergetwosortedlist

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	h1 := &ListNode{Val: 1}
	h1.Add(3).Add(5)
	h2 := &ListNode{Val: 2}
	h2.Add(4).Add(6)

	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{h1, h2}, want: "123456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2).Show(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
