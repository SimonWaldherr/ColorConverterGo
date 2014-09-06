package colorconverter

import (
	"testing"
)

func Test_RGB2HSL(t *testing.T) {
	h, s, l := RGB2HSL(121, 167, 22)

	if (h != 79) || (s != 76) || (l != 37) {
		t.Error("wrong values returned")
	}
}

func Test_HSL2RGB(t *testing.T) {
	r, g, b := HSL2RGB(350, 30, 50)

	if (r != 165) || (g != 89) || (b != 102) {
		t.Error("wrong values returned")
	}
}

func Test_RGB2CMYK(t *testing.T) {
	c, m, y, k := RGB2CMYK(10, 20, 200)

	if (c != 242) || (m != 229) || (y != 0) || (k != 55) {
		t.Error("wrong values returned")
	}
}

func Test_HSV2RGB(t *testing.T) {
	r, g, b := HSV2RGB(184, 71, 81)

	if (r != 59) || (g != 196) || (b != 206) {
		t.Error("wrong values returned")
	}
}

func Test_HEX2RGB(t *testing.T) {
	r, g, b := HEX2RGB("#C3C3FF")

	if (r != 195) || (g != 195) || (b != 255) {
		t.Error("wrong values returned")
	}
}


func Test_HEX2X2HEX(t *testing.T) {
	in := "abcdef"
	out := RGB2HEX(HEX2RGB(in))
	//out := RGB2HEX(CMYK2RGB(HSL2CMYK(RGB2HSL(HEX2RGB(in)))))
	if in != out {
		t.Error("wrong values returned:" + in + " != " + out)
	}
}
