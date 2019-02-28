package main

import (
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	result := printHelloWorld()
	if result != "Hello World." {
		t.Fatal("Failed.")
	}
}
