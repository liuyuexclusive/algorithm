package tree

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{args: args{root}, want: [][]int{
			{3},
			{9, 20},
			{15, 7},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestCodec_serialize(t *testing.T) {
// 	type args struct {
// 		root *TreeNode
// 	}
// 	root := &TreeNode{Val: 3}
// 	root.Left = &TreeNode{Val: 9}
// 	root.Right = &TreeNode{Val: 20}
// 	root.Right.Left = &TreeNode{Val: 15}
// 	root.Right.Right = &TreeNode{Val: 7}
// 	tests := []struct {
// 		name string
// 		this *Codec
// 		args args
// 		want string
// 	}{
// 		{args: args{root: root}, want: "3,9,20,N,N,15,7,N,N,N,N"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			this := &Codec{}
// 			if got := this.serialize(tt.args.root); got != tt.want {
// 				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestCodec_deserialize(t *testing.T) {
// 	type args struct {
// 		data string
// 	}

// 	root := &TreeNode{Val: 3}
// 	root.Left = &TreeNode{Val: 9}
// 	root.Right = &TreeNode{Val: 20}
// 	root.Right.Left = &TreeNode{Val: 15}
// 	root.Right.Right = &TreeNode{Val: 7}
// 	tests := []struct {
// 		name string
// 		this *Codec
// 		args args
// 		want *TreeNode
// 	}{
// 		{args: args{data: ""}, want: nil},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			this := &Codec{}
// 			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_findWords(t *testing.T) {
// 	type args struct {
// 		board [][]byte
// 		words []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []string
// 	}{
// 		// {args: args{board: [][]byte{
// 		// 	{'o', 'a', 'a', 'n'},
// 		// 	{'e', 't', 'a', 'e'},
// 		// 	{'i', 'h', 'k', 'r'},
// 		// 	{'i', 'f', 'l', 'v'},
// 		// }, words: []string{"oath"}}, want: []string{"oath"}},
// 		{args: args{board: [][]byte{
// 			{'o', 'a', 'a', 'n'},
// 			{'e', 't', 'a', 'e'},
// 			{'i', 'h', 'k', 'r'},
// 			{'i', 'f', 'l', 'v'},
// 		}, words: []string{"oath", "pea", "eat", "rain"}}, want: []string{"eat", "oath"}},
// 		{args: args{board: [][]byte{{'a', 'a'}}, words: []string{"aaa"}}, want: []string{}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("findWords() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
