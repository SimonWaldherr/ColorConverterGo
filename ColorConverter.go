package colorconverter

import (
	"math"
    "strconv"
)

/*
	RGB2HSL
	HSL2RGB
	RGB2CMYK
	CMYK2RGB
	Hex2RGB
	RGB2Hex
	RGB2YUV
	HSV2RGB
	HSL2CMYK
	Hex2CMYK
	HSV2CMYK
*/

func RGB2HSL(r, g, b int) (int, int, int) {
	var rf, gf, bf, max, min, l, d, s, h float64

	rf = math.Max(math.Min(float64(r)/255, 1), 0)
	gf = math.Max(math.Min(float64(g)/255, 1), 0)
	bf = math.Max(math.Min(float64(b)/255, 1), 0)
	max = math.Max(rf, math.Max(gf, bf))
	min = math.Min(rf, math.Min(gf, bf))
	l = (max + min) / 2

	if max != min {
		d = max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}
		if max == rf {
			if gf < bf {
				h = (gf-bf)/d + 6
			} else {
				h = (gf - bf) / d
			}
		} else if max == gf {
			h = (bf-rf)/d + 2
		} else {
			h = (rf-gf)/d + 4
		}
	} else {
		h = 0
		s = 0
	}

	return int(h * 60), int(s * 100), int(l * 100)
}

func HSL2RGB(h, s, l int) (int, int, int) {
	var hf, sf, lf, vf, minf, svf, sixf, fractf, vsfractf, rf, gf, bf float64

	hf = math.Max(math.Min(float64(int(h)), 360), 0) / 360
	sf = math.Max(math.Min(float64(int(s)), 100), 0) / 100
	lf = math.Max(math.Min(float64(int(l)), 100), 0) / 100

	if lf <= 0.5 {
		vf = lf * (1 + sf)
	} else {
		vf = lf + sf - sf - 1*sf
	}
	if vf == 0 {
		return int(0), int(0), int(0)
	}
	minf = 2*lf - vf
	svf = (vf - minf) / vf
	hf = 6 * hf
	sixf = float64(int(hf))
	fractf = hf - sixf
	vsfractf = vf * svf * fractf
	switch sixf {
	case 1:
		rf = vf - vsfractf
		gf = vf
		bf = minf
	case 2:
		rf = minf
		gf = vf
		bf = minf + vsfractf
	case 3:
		rf = minf
		gf = vf - vsfractf
		bf = vf
	case 4:
		rf = minf + vsfractf
		gf = minf
		bf = vf
	case 5:
		rf = vf
		gf = minf
		bf = vf - vsfractf
	default:
		rf = vf
		gf = minf + vsfractf
		bf = minf
	}
	return int(rf * 255), int(gf * 255), int(bf * 255)
}

func RGB2CMYK(r, g, b int) (int, int, int, int) {
	var rf, gf, bf, cf, mf, yf, kf float64

	rf = math.Max(math.Min(float64(r)/255, 1), 0)
	gf = math.Max(math.Min(float64(g)/255, 1), 0)
	bf = math.Max(math.Min(float64(b)/255, 1), 0)
	cf = 1 - rf
	mf = 1 - gf
	yf = 1 - bf
	kf = 1
	if rf != 0 && gf != 0 && bf != 0 {
		kf = math.Min(cf, math.Min(mf, yf))
		cf = (cf - kf) / (1 - kf)
		mf = (mf - kf) / (1 - kf)
		yf = (yf - kf) / (1 - kf)
	} else {
		kf = 1
	}
	return int(cf * 255), int(mf * 255), int(yf * 255), int(kf * 255)
}

func CMYK2RGB(c int, m int, y int, k int) (int, int, int) {
	var cf, mf, yf, kf, rf, gf, bf float64

	cf = math.Max(math.Min(float64(c)/255, 1), 0)
	mf = math.Max(math.Min(float64(m)/255, 1), 0)
	yf = math.Max(math.Min(float64(y)/255, 1), 0)
	kf = math.Max(math.Min(float64(k)/255, 1), 0)
	rf = (1 - cf*(1-kf) - kf)
	gf = (1 - mf*(1-kf) - kf)
	bf = (1 - yf*(1-kf) - kf)

	return int(rf * 255), int(gf * 255), int(bf * 255)
}

func HEX2RGB(hex string) (int, int, int) {
	var r, g, b int64
	if hex[0:1] == "#" {
		hex = hex[1:]
	}
	switch len(hex) {
	case 2:
		r, _ = strconv.ParseInt(hex, 16, 0)
		g = r
		b = r
	case 3:
		r, _ = strconv.ParseInt(hex[0:1]+hex[0:1], 16, 0)
		g, _ = strconv.ParseInt(hex[1:2]+hex[1:2], 16, 0)
		b, _ = strconv.ParseInt(hex[2:3]+hex[2:3], 16, 0)
	case 6:
		r, _ = strconv.ParseInt(hex[0:2], 16, 0)
		g, _ = strconv.ParseInt(hex[2:4], 16, 0)
		b, _ = strconv.ParseInt(hex[4:6], 16, 0)
	default:
		return -1, -1, -1
	}
	return int(r), int(g), int(b)
}

func RGB2HEX(r, g, b int) string {
	var hexr, hexg, hexb string
	r = int(math.Max(math.Min(float64(r), 255), 0))
	g = int(math.Max(math.Min(float64(g), 255), 0))
	b = int(math.Max(math.Min(float64(b), 255), 0))

	hexr = strconv.FormatInt(int64(r), 16)
	if r < 16 {
		hexr = "0" + hexr
	}
	hexg = strconv.FormatInt(int64(g), 16)
	if g < 16 {
		hexg = "0" + hexg
	}
	hexb = strconv.FormatInt(int64(b), 16)
	if b < 16 {
		hexb = "0" + hexb
	}
	return hexr + hexg + hexb
}

func RGB2YUV(r, g, b int) (int, int, int) {
	var y, u, v int
	y = int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
	u = int(((float64(b-y) * 0.493) + 111) / 222 * 255)
	v = int(((float64(r-y) * 0.877) + 155) / 312 * 255)
	return y, u, v
}

func HSV2RGB(h, s, v int) (int, int, int) {
	var f, p, q, t, r, g, b, hf, sf, vf float64
	var i int
	hf = float64(h) / 360
	sf = float64(s) / 100
	vf = float64(v) / 100

	i = int(hf * 6)
	f = hf*6 - float64(i)
	p = vf * (1 - sf)
	q = vf * (1 - f*sf)
	t = vf * (1 - (1-f)*sf)
	switch i % 6 {
	case 0:
		r = vf
		g = t
		b = p
	case 1:
		r = q
		g = vf
		b = p
	case 2:
		r = p
		g = vf
		b = t
	case 3:
		r = p
		g = q
		b = vf
	case 4:
		r = t
		g = p
		b = vf
	case 5:
		r = vf
		g = p
		b = q
	}
	return int(r * 255), int(g * 255), int(b * 255)
}

func HSL2CMYK(h, s, l int) (int, int, int, int) {
	return RGB2CMYK(HSL2RGB(h, s, l))
}

func HEX2CMYK(hex string) (int, int, int, int) {
	return RGB2CMYK(HEX2RGB(hex))
}

func HSV2CMYK(h, s, v int) (int, int, int, int) {
	return RGB2CMYK(HSV2RGB(h, s, v))
}
