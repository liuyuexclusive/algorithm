package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.add(2).add(3).add(4).add(5)

	rev := reverseList(head)
	got := rev.show()
	want := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func Test_hasCycle(t *testing.T) {
	h1 := &ListNode{Val: 1}
	h1.add(2).add(3).Next = h1
	h2 := &ListNode{Val: 1}
	h2.add(2).add(3).add(1)
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args: args{head: h1}, want: true},
		{args: args{head: h2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l1.add(3).add(5).add(7).add(9)
	l2.add(4).add(6).add(8).add(10)
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{args: args{l1: l1, l2: l2}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2).show(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	h1 := &ListNode{Val: 1}
	h1.add(4).add(7)
	h2 := &ListNode{Val: 2}
	h2.add(5)
	h3 := &ListNode{Val: 3}
	h3.add(6)
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{args: args{[]*ListNode{h1, h2, h3}}, want: []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists).show(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	h1 := &ListNode{Val: 1}
	h1.add(2).add(3).add(4).add(5)
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{args: args{head: h1, n: 2}, want: []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n).show(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reorderList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.add(2).add(3).add(4)
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{args: args{head}, want: []int{1, 4, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if got := tt.args.head.show(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
