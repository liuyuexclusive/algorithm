package circle

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := &ListNode{Val: 1}
	head1.Next = head1
	head2 := &ListNode{Val: 1}
	head2.Add(1)
	head3 := &ListNode{Val: 1}
	head3.Add(2).Add(3).Add(4).Next = head3
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{head1}, want: true},
		{args: args{head2}, want: false},
		{args: args{head3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
