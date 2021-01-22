package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	n1 := New(-8)
	n2 := New(-10)
	n2.Add(-6)
	n3 := New(-10)
	n3.Add(-9)
	n3.Add(-8)
	n4 := New(-10)
	slice := []*ListNode{
		n1,
		n2,
		n3,
		n4,
	}

	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{args: args{lists: slice}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedianFinder(t *testing.T) {
	c := Constructor()
	c.AddNum(1)
	c.AddNum(2)

	fmt.Println(c.FindMedian())
	c.AddNum(3)
	fmt.Println(c.FindMedian())
}
