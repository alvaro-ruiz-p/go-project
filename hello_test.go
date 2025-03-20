package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Alvaro")
	want := "hello, Alvaro"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
