package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 1

	if a != b {
		t.Fatal("FAILED")
	}
}

