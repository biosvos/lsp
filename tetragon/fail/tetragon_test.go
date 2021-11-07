package fail

import (
	"strconv"
	"testing"
)

func TestTetragonal(t *testing.T) {
	var tetragons []Tetragon

	tetragons = append(tetragons, NewRectangle(1, 1))
	tetragons = append(tetragons, NewSquare(1))

	for i, tetragon := range tetragons {
		t.Run(strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			tetragon.SetWidth(5)
			tetragon.SetHeight(2)
			area := tetragon.Area()
			if area != 10 {
				t.Fatalf("%v is not 10", area)
			}
		})
	}
}
