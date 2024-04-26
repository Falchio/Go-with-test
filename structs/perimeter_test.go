package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertEquals(t, want, got)
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		got := rectangle.Area()
		want := 72.0
		assertEquals(t, want, got)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		assertEquals(t, want, got)
	})

}

func assertEquals(t *testing.T, want float64, got float64) {
	t.Helper()
	if got != want {
		t.Errorf("want %g but got %g", want, got)
	}
}
