package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world!"
	get := Hello()
	if want != get {
		t.Errorf("want %q get %q", want, get)
	}
}
