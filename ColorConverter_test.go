package colorconverter

import (
	"testing"
)

func Test_RGB2HSL(t *testing.T) {
	r, g, b := 121, 167, 22
	h, s, l := RGB2HSL(r, g, b)

	t.Logf("RGB(%v, %v, %v) => HSL(%v, %v, %v)\n", r, g, b, h, s, l)
	if (h != 79) || (s != 76) || (l != 37) {
		t.Error("wrong values returned")
	}
}

func Test_RGB2YUV(t *testing.T) {
	r, g, b := 121, 167, 22
	y, u, v := RGB2YUV(r, g, b)

	t.Logf("RGB(%v, %v, %v) => YUV(%v, %v, %v)\n", r, g, b, y, u, v)
	if (y != 136) || (u != 62) || (v != 115) {
		t.Error("wrong values returned")
	}
}

func Test_HSL2RGB(t *testing.T) {
	h, s, l := 350, 30, 50
	r, g, b := HSL2RGB(h, s, l)

	t.Logf("HSL(%v, %v, %v) => RGB(%v, %v, %v)\n", h, s, l, r, g, b)
	if (r != 165) || (g != 89) || (b != 102) {
		t.Error("wrong values returned")
	}
}

func Test_HSL2RGB2(t *testing.T) {
	h, s, l := 15, 30, 50
	r, g, b := HSL2RGB(h, s, l)

	t.Logf("HSL(%v, %v, %v) => RGB(%v, %v, %v)\n", h, s, l, r, g, b)
	if (r != 165) || (g != 108) || (b != 89) {
		t.Error("wrong values returned")
	}
}

func Test_HSL2RGB3(t *testing.T) {
	h, s, l := 50, 50, 50
	r, g, b := HSL2RGB(h, s, l)

	t.Logf("HSL(%v, %v, %v) => RGB(%v, %v, %v)\n", h, s, l, r, g, b)
	if (r != 191) || (g != 170) || (b != 63) {
		t.Error("wrong values returned")
	}
}

func Test_HSV2CMYK(t *testing.T) {
	h, s, v := 10, 20, 10
	c, m, y, k := HSV2CMYK(h, s, v)

	t.Logf("HSV(%v, %v, %v) => CMYK(%v, %v, %v, %v)\n", h, s, v, c, m, y, k)
	if (c != 0) || (m != 40) || (y != 51) || (k != 230) {
		t.Error("wrong values returned")
	}
}

func Test_RGB2CMYK(t *testing.T) {
	r, g, b := 10, 20, 200
	c, m, y, k := RGB2CMYK(r, g, b)

	t.Logf("RGB(%v, %v, %v) => CMYK(%v, %v, %v, %v)\n", r, g, b, c, m, y, k)
	if (c != 242) || (m != 229) || (y != 0) || (k != 55) {
		t.Error("wrong values returned")
	}
}

func Test_CMYK2RGB(t *testing.T) {
	c, m, y, k := 242, 229, 0, 55
	r, g, b := CMYK2RGB(c, m, y, k)

	t.Logf("CMYK(%v, %v, %v, %v) => RGB(%v, %v, %v)\n", c, m, y, k, r, g, b)
	if (r != 10) || (g != 20) || (b != 200) {
		t.Error("wrong values returned")
	}
}

func Test_HSV2RGB(t *testing.T) {
	h, s, v := 184, 71, 81
	r, g, b := HSV2RGB(h, s, v)

	t.Logf("HSV(%v, %v, %v) => RGB(%v, %v, %v)\n", h, s, v, r, g, b)
	if (r != 59) || (g != 196) || (b != 206) {
		t.Error("wrong values returned")
	}
}

func Test_HEX2RGB(t *testing.T) {
	hex := "#C3C3FF"
	r, g, b := HEX2RGB(hex)

	t.Logf("%v => RGB(%v, %v, %v)\n", hex, r, g, b)
	if (r != 195) || (g != 195) || (b != 255) {
		t.Error("wrong values returned")
	}
}

func Test_HEX2HSL(t *testing.T) {
	hex := "#C3C3FF"
	r, g, b := HEX2RGB(hex)
	h, s, l := RGB2HSL(r, g, b)

	t.Logf("%v => HSL(%v, %v, %v)\n", hex, h, s, l)
	if (h != 240) || (s != 100) || (l != 88) {
		t.Error("wrong values returned")
	}
}

func Test_HEX2HSL2(t *testing.T) {
	hex := "#0ff0aa"
	r, g, b := HEX2RGB(hex)
	h, s, l := RGB2HSL(r, g, b)

	t.Logf("%v => HSL(%v, %v, %v)\n", hex, h, s, l)
	if (h != 161) || (s != 88) || (l != 50) {
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

	in = "f9f9fa"
	out = RGB2HEX(HEX2RGB(in))
	if in != out {
		t.Error("wrong values returned:" + in + " != " + out)
	}
}

func Test_HEX2X2HEX2(t *testing.T) {
	in := "fff"
	out := RGB2HEX(HEX2RGB(in))
	if "ffffff" != out {
		t.Error("wrong values returned:" + in + " != " + out)
	}

	in = "123"
	out = RGB2HEX(HEX2RGB(in))
	if "112233" != out {
		t.Error("wrong values returned:" + in + " != " + out)
	}

	in = "fe"
	out = RGB2HEX(HEX2RGB(in))
	if "fefefe" != out {
		t.Error("wrong values returned:" + in + " != " + out)
	}

	in = "0a"
	out = RGB2HEX(HEX2RGB(in))
	if "0a0a0a" != out {
		t.Error("wrong values returned:" + in + " != " + out)
	}
}

func Test_WrongHEX(t *testing.T) {
	in := "f"
	r, g, b := HEX2RGB(in)
	if (r != -1) || (g != -1) || (b != -1) {
		t.Error("wrong values returned")
	}
}
