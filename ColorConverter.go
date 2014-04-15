package main

import (
	"fmt"
	"math"
)

func RGB2HSL(r int, g int, b int) (int, int, int) {
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

func HSV2RGB(h int, s int, v int) (int, int, int) {
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

func main() {
	fmt.Println(RGB2HSL(121, 167, 22))
	fmt.Println(RGB2HSL(69, 209, 237))
	fmt.Println(RGB2HSL(254, 207, 37))
	fmt.Println(RGB2HSL(122, 167, 255))
	fmt.Println(RGB2HSL(255, 255, 255))
	fmt.Println(HSV2RGB(184, 71, 81))

	fmt.Println(RGB2HSL(HSV2RGB(184, 71, 81)))
	fmt.Println(RGB2HSL(CMYK2RGB(215, 55, 0, 20)))
}
