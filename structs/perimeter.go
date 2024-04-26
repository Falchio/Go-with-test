package structs

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
