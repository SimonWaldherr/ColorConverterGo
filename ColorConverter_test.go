package colorconverter

import (
	"github.com/SimonWaldherr/ColorConverterGo"
	"testing"
)

func Test_RGB2HSL(t *testing.T) {
	h, s, l := colorconverter.RGB2HSL(121, 167, 22)

	if (h != 79) || (s != 76) || (l != 37) {
		t.Error("wrong values returned")
	}
}

func Test_HSL2RGB(t *testing.T) {
	r, g, b := colorconverter.HSL2RGB(350, 30, 50)

	if (r != 165) || (g != 89) || (b != 102) {
		t.Error("wrong values returned")
	}
}
