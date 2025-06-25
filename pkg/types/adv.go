package types

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"

	c "github.com/fopina/zengge-led-ctl/pkg/constants"
	"github.com/go-ble/ble"
)

type ZenggeAdvertisementDetails struct {
	Firmware    byte
	MAC         net.HardwareAddr
	On          bool
	Mode        byte
	Brightness  uint16
	RGB         []byte
	Temperature uint8
	LEDCount    uint8
}

func NewZenggeAdvertisementDetails(data []byte) *ZenggeAdvertisementDetails {
	if len(data) != 29 {
		return nil
	}
	// per https://github.com/8none1/zengge_lednetwf/blob/main/readme.md#advertising-data
	d := &ZenggeAdvertisementDetails{
		Firmware:    data[2],
		MAC:         data[4:10],
		On:          data[16] == c.PowerOnByte,
		Mode:        data[17],
		Brightness:  binary.LittleEndian.Uint16(data[18:20]),
		RGB:         data[20:23],
		Temperature: data[23],
		LEDCount:    data[26],
	}
	return d
}

func (z ZenggeAdvertisementDetails) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("MAC: %s ", z.MAC))
	if z.On {
		b.WriteString("ON ")
	} else {
		b.WriteString("OFF ")
	}
	b.WriteString(fmt.Sprintf("Mode: %x Brightness: %d RGB: %v Temperature: %d LEDs: %d", z.Mode, z.Brightness, z.RGB, z.Temperature, z.LEDCount))
	return b.String()
}

// ZenggeAdvertisement ...
type ZenggeAdvertisement struct {
	Name        string
	Addr        ble.Addr
	Connectable bool
	RSSI        int
	MD          []byte
	Details     *ZenggeAdvertisementDetails
}

func (z ZenggeAdvertisement) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("[%s] ", z.Addr))
	if z.Connectable {
		b.WriteString("C ")
	} else {
		b.WriteString("N ")
	}
	b.WriteString(fmt.Sprintf("%3d: Name %s", z.RSSI, z.Name))
	if len(z.MD) > 0 {
		if z.Details == nil {
			b.WriteString(fmt.Sprintf(" [MD: %X]", z.MD))
		} else {
			b.WriteString(fmt.Sprintf(" [%s]", z.Details))
		}
	}
	return b.String()
}
