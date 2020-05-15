package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTraverse(t *testing.T) {
	tr := New()
	var want []Value
	for i := 1; i < 16; i++ {
		val := fmt.Sprintf("Scarlett%d", i)
		tr.Put(i, val)
		want = append(want, val)
	}
	res := tr.Traverse()
	if !reflect.DeepEqual(want, res) {
		t.Errorf("want:%v,get:%v", want, res)
	}
}

func TestDeepth(t *testing.T) {
	tr := New()
	for i := 1; i < 16; i++ {
		val := fmt.Sprintf("Scarlett%d", i)
		tr.Put(i, val)
	}
	want := 4
	res := tr.Deepth()
	if !reflect.DeepEqual(want, res) {
		t.Errorf("want:%v,get:%v", want, res)
	}
}
