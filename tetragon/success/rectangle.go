package fail

type Rectangle struct {
	height, width uint64
}

func NewRectangle(height uint64, width uint64) *Rectangle {
	return &Rectangle{height: height, width: width}
}

func (r *Rectangle) SetHeight(height uint64) {
	r.height = height
}

func (r *Rectangle) SetWidth(width uint64) {
	r.width = width
}

func (r *Rectangle) Area() uint64 {
	return r.width * r.height
}
