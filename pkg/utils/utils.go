package utils

import "math"

// RGBToHSV Convert RGB code to HSV - provided by ChatGPT
func RGBToHSV(r, g, b uint8) (h, s, v float64) {
	// Normalize RGB to [0, 1]
	fr := float64(r) / 255.0
	fg := float64(g) / 255.0
	fb := float64(b) / 255.0

	max := math.Max(fr, math.Max(fg, fb))
	min := math.Min(fr, math.Min(fg, fb))
	delta := max - min

	// Hue
	switch {
	case delta == 0:
		h = 0
	case max == fr:
		h = 60 * math.Mod((fg-fb)/delta, 6)
	case max == fg:
		h = 60 * ((fb-fr)/delta + 2)
	case max == fb:
		h = 60 * ((fr-fg)/delta + 4)
	}
	if h < 0 {
		h += 360
	}

	// Saturation
	if max == 0 {
		s = 0
	} else {
		s = delta / max
	}

	// Value
	v = max

	return
}

// RGBToHSV_bytes Convert RGB code to HSV, truncated to bytes
func RGBToHSV_bytes(r, g, b uint8) (h, s, v byte) {
	hf, sf, vf := RGBToHSV(r, g, b)
	h = byte(int(hf) / 2)
	s = byte(sf * 100)
	v = byte(vf * 100)
	return
}
