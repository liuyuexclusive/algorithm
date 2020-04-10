package reorderlist

import (
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	h1 := &ListNode{Val: 1}
	h1.Add(2).Add(3).Add(4).Add(5)
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{head: h1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
	t.Log(h1.Show())
}
