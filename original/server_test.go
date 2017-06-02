package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	value := 1
	expected := 1
	if value != expected {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}
