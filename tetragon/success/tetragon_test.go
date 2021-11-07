package fail

import (
	"strconv"
	"testing"
)

func TestTetragonal(t *testing.T) {
	var tetragons []Tetragon

	tetragons = append(tetragons, NewRectangle(5, 5))
	tetragons = append(tetragons, NewSquare(5))

	for i, tetragon := range tetragons {
		t.Run(strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			area := tetragon.Area()
			if area != 25 {
				t.Fatalf("%v is not 25", area)
			}
		})
	}
}
