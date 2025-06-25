package types

import (
	"fmt"

	"github.com/fopina/zengge-led-ctl/pkg/utils"
)

// RGBColor ...
type RGBColor struct {
	Red,
	Green,
	Blue byte
}

func (c RGBColor) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.Red, c.Green, c.Blue)
}

func (c RGBColor) ConvertToHSV() HSVColor {
	h, s, v := utils.RGBToHSV_bytes(c.Red, c.Green, c.Blue)
	return HSVColor{Hue: h, Saturation: s, Value: v}
}

// HSVColor ...
type HSVColor struct {
	Hue,
	Saturation,
	Value byte
}

func (c HSVColor) String() string {
	return fmt.Sprintf("hsv(%d, %d, %d)", c.Hue, c.Saturation, c.Value)
}
