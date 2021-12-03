package main

import "testing"

func TestInitMin(t *testing.T) {
	ans := InitMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}
