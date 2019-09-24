package shape

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectanle := Rectangle{12.0, 6.0}
		checkArea(t, rectanle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{12, 6}, hasArea: 72.0},
		{shape: Circle{10}, hasArea: 314.1592653589793},
		{shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#vgot %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
