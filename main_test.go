package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Johnny")
	want := "Hello, Johnny"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
