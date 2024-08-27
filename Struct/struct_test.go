package structure

import "testing"

func StructTest (t *testing.T) {
	rectangle := Rect{10, 6}
	got := Perimeter(rectangle)
	want := 32.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}