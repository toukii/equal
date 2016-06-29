package point

import (
	. "github.com/toukii/equal"
	"testing"
)

func TestAsserts(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{2, 1}

	Equal(t, p1, p2)
}
