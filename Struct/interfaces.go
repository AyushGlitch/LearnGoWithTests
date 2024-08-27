package structure

type Rectangle struct {
	Width	float64
	Length	float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Shape interface {
	Area()	float64
}