package store

import "testing"

func TestAdd(t *testing.T) {
	s := &Store{}
	s.Add("x")
	if len(s.items) != 1 {
		t.Fatal("nope")
	}
}
