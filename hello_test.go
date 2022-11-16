package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris"

	if got != want {
		t.Errorf("got %q got %q", got, want)
	}
}
