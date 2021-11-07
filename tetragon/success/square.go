package fail

type Square struct {
	side uint64
}

func NewSquare(side uint64) *Square {
	return &Square{side: side}
}

func (s *Square) SetSide(side uint64) {
	s.side = side
}

func (s *Square) Area() uint64 {
	return s.side * s.side
}
