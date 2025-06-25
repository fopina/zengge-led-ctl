package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/dev"
	"github.com/fopina/zengge-led-ctl/pkg/utils"
	"github.com/go-ble/ble"
	"github.com/pkg/errors"
)

type ZenggeClient struct {
	deviceName    string
	device        ble.Device
	client        ble.Client
	notifyChar    *ble.Characteristic
	writeChar     *ble.Characteristic
	packetCounter uint16
}

func NewZenggeClient(device string) (*ZenggeClient, error) {
	d, err := dev.NewDevice(device)
	if err != nil {
		return nil, fmt.Errorf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)
	return &ZenggeClient{
		deviceName:    device,
		device:        d,
		packetCounter: 0,
	}, nil
}

func (c *ZenggeClient) Scan(duration time.Duration, duplicates bool, handler ScanHandler) error {
	scanHandler := func(a ble.Advertisement) {
		if !strings.HasPrefix(a.LocalName(), "LEDnetWF") {
			return
		}
		adv := ZenggeAdvertisement{
			Name:        a.LocalName(),
			Addr:        a.Addr(),
			Connectable: a.Connectable(),
			RSSI:        a.RSSI(),
			MD:          a.ManufacturerData(),
			Details:     NewZenggeAdvertisementDetails(a.ManufacturerData()),
		}
		handler(adv)
	}

	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), duration))
	return chkScanErr(ble.Scan(ctx, duplicates, scanHandler, nil))
}

func (c *ZenggeClient) Connect(addr string, duration time.Duration) error {
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), duration))
	// Dial directly, skip Connect - it's slower (as it scans) and DOESN'T EVEN WORK!... https://github.com/go-ble/ble/pull/112
	client, err := ble.Dial(ctx, ble.NewAddr(addr))
	if err != nil {
		return err
	}
	c.client = client

	_, err = client.DiscoverProfile(true)
	if err != nil {
		return fmt.Errorf("can't discover profile: %s", err)
	}

	notify := client.Profile().FindCharacteristic(ble.NewCharacteristic(ble.MustParse(NotifyUUID)))
	if notify == nil {
		return fmt.Errorf("cannot find characteristic to subscribe")
	}
	c.notifyChar = notify

	write := client.Profile().FindCharacteristic(ble.NewCharacteristic(ble.MustParse(WriteUUID)))
	if write == nil {
		return fmt.Errorf("cannot find characteristic to write")
	}
	c.writeChar = write
	return nil
}

func (c *ZenggeClient) Subscribe(handler ble.NotificationHandler) error {
	notificationhandler := func(req []byte) {
		// TODO: implement parsing
		handler(req)
	}

	return c.client.Subscribe(c.notifyChar, false, notificationhandler)
}

// SendInitialPackage ??? TBD what this is and when is it required.....
func (c *ZenggeClient) SendInitialPacket() error {
	InitialPacket[0] = 0x0
	return c.client.WriteCharacteristic(c.writeChar, c.preparePacket(InitialPacket), false)
}

// GetStripSettings ??? TBD what this is...
func (c *ZenggeClient) GetStripSettings() error {
	return c.client.WriteCharacteristic(c.writeChar, c.preparePacket(GetStripSettingsPacket), false)
}

// preparePacket updates the counter bytes (first two) of a packet IN PLACE and returns the same reference
func (c *ZenggeClient) preparePacket(packet []byte) []byte {
	// counter seems to be ignored but let's pretend it is not
	c.packetCounter++
	packet[0] = byte(c.packetCounter >> 8)
	packet[1] = byte(c.packetCounter)
	return packet
}

// PowerOff Power off the LED strip
func (c *ZenggeClient) PowerOff() error {
	PowerPacket[9] = byte(PowerOffByte)
	return c.client.WriteCharacteristic(c.writeChar, c.preparePacket(PowerPacket), false)
}

// PowerOn Power on the LED strip
func (c *ZenggeClient) PowerOn() error {
	PowerPacket[9] = byte(PowerOnByte)
	return c.client.WriteCharacteristic(c.writeChar, c.preparePacket(PowerPacket), false)
}

// SetWhite Set LED color to white
func (c *ZenggeClient) SetWhite() error {
	return c.client.WriteCharacteristic(c.writeChar, c.preparePacket(WhitePacket), false)
}

// SetRGB Set LED color to color defined by RGB
func (c *ZenggeClient) SetRGB(red byte, green byte, blue byte) error {
	packet := c.preparePacket(HsvPacket)
	h, s, v := utils.RGBToHSV_bytes(red, green, blue)
	packet[10] = h
	packet[11] = s
	packet[12] = v
	return c.client.WriteCharacteristic(c.writeChar, packet, false)
}

func chkScanErr(err error) error {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		return err
	}
	return nil
}
