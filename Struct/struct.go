package structure

type Rect struct {
	Width	float64
	Length	float64
}

func Perimeter (rectangle Rect) float64 {
	return 2 * (rectangle.Width + rectangle.Length)
}