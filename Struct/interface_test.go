package structure

import "testing"


func InterfaceTest (t *testing.T) {
	checkArea := func (t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 10}
		checkArea(t, rectangle, 120)
	})
}