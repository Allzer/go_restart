package rectangle

type Rectangle struct {
	A int
	B int
}

func (r Rectangle) Area() float32 {
	return float32(r.A * r.B)
}
