package models

import "testing"

func TestSum(t *testing.T) {
	r := Sum(2, 3)
	if r != 5 {
		t.Errorf("expected 5; got %d", r)
	}
}

func Sum(a, b int) int {
	return a + b
}
