package fail

type Square struct {
	side uint64
}

func NewSquare(side uint64) *Square {
	return &Square{side: side}
}

func (s *Square) SetHeight(height uint64) {
	s.side = height
}

func (s *Square) SetWidth(width uint64) {
	s.side = width
}

func (s *Square) Area() uint64 {
	return s.side * s.side
}
