package client

import (
	"fmt"
	"strings"

	"github.com/go-ble/ble"
)

// ZenggeAdvertisement ...
type ZenggeAdvertisement struct {
	Name        string
	Addr        ble.Addr
	Connectable bool
	RSSI        int
	MD          []byte
}

// ScanHandler handles Zengge advertisements.
type ScanHandler func(a ZenggeAdvertisement)

func (z ZenggeAdvertisement) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("[%s] ", z.Addr))
	if z.Connectable {
		b.WriteString("C ")
	} else {
		b.WriteString("N ")
	}
	b.WriteString(fmt.Sprintf("%3d: Name %s, MD: %X", z.RSSI, z.Name, z.MD))
	return b.String()
}
