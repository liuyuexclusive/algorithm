package removefromend

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	h1 := &ListNode{Val: 1}
	h1.Add(2).Add(3).Add(4).Add(5)
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{h1, 2}, want: "1235"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n).Show(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
