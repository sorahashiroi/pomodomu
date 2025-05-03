package main

import "testing"

func Example_pomodomu() {
	goMain([]string{"pomodomu"})
}

func TestHello(t *testing.T) {
	got := hello()
	want := "Welcome to Pomodomu!"
	if got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
