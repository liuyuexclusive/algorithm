package linkedList

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	tests := []struct {
		name string
		args args
	}{
		{"", args{n1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
