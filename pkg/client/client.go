package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/dev"
	"github.com/go-ble/ble"
	"github.com/pkg/errors"
)

type ZenggeClient struct {
	deviceName string
	device     ble.Device
}

func NewZenggeClient(device string) (*ZenggeClient, error) {
	d, err := dev.NewDevice(device)
	if err != nil {
		return nil, fmt.Errorf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)
	return &ZenggeClient{
		deviceName: device,
		device:     d,
	}, nil
}

func (c *ZenggeClient) Scan(duration time.Duration, duplicates bool, handler ble.AdvHandler) error {
	scanHandler := func(a ble.Advertisement) {
		if !strings.HasPrefix(a.LocalName(), "LEDnetWF") {
			return
		}
		handler(a)
	}

	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), duration))
	return chkScanErr(ble.Scan(ctx, duplicates, scanHandler, nil))
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
