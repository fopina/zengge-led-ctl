package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRGBToHSV(t *testing.T) {
	testCases := []struct {
		name     string
		r        uint8
		g        uint8
		b        uint8
		expected []float64
	}{
		{
			name:     "greenish",
			r:        3,
			g:        252,
			b:        102,
			expected: []float64{143.85542168674698, 0.9880952380952381, 0.9882352941176471},
		},
		{
			name:     "reddish",
			r:        252,
			g:        3,
			b:        48,
			expected: []float64{349.1566265060241, 0.9880952380952381, 0.9882352941176471},
		},
		{
			name:     "darkred",
			r:        28,
			g:        9,
			b:        9,
			expected: []float64{0, 0.6785714285714286, 0.10980392156862745},
		},
		{
			name:     "red",
			r:        201,
			g:        62,
			b:        62,
			expected: []float64{0, 0.6915422885572139, 0.788235294117647},
		},
	}

	for _, tc := range testCases {
		h, s, v := RGBToHSV(tc.r, tc.g, tc.b)

		assert.Equal(t, tc.expected, []float64{h, s, v}, tc.name)
	}
}

func TestRGBToHSV_bytes(t *testing.T) {
	testCases := []struct {
		name     string
		r        uint8
		g        uint8
		b        uint8
		expected []byte
	}{
		{
			name:     "greenish",
			r:        3,
			g:        252,
			b:        102,
			expected: []byte{0x47, 0x62, 0x62},
		},
		{
			name:     "reddish",
			r:        252,
			g:        3,
			b:        48,
			expected: []byte{0xae, 0x62, 0x62},
		},
		{
			name:     "darkred",
			r:        28,
			g:        9,
			b:        9,
			expected: []byte{0x0, 0x43, 0xa},
		},
		{
			name:     "red",
			r:        201,
			g:        62,
			b:        62,
			expected: []byte{0x0, 0x45, 0x4e},
		},
	}

	for _, tc := range testCases {
		h, s, v := RGBToHSV_bytes(tc.r, tc.g, tc.b)

		assert.Equal(t, tc.expected, []byte{h, s, v}, tc.name)
	}
}
