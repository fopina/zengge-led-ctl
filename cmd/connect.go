// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/dev"
	"github.com/go-ble/ble"
	"github.com/spf13/cobra"
)

// all taken straight from https://github.com/8none1/zengge_lednetwf/blob/f72474c502ed15767b4fbbe8ef028fee980b9b11/ledwf_controller.py#L13
const ServiceUUID = "0000ffff-0000-1000-8000-00805f9b34fb"

// const NotifyUUID = "0000ff02-0000-1000-8000-00805f9b34fb"
const NotifyUUID = "ff02"

// const WriteUUID = "0000ff01-0000-1000-8000-00805f9b34fb"
const WriteUUID = "ff01"
const PixelCount = 48

var InitialPacket = []byte{0x00, 0x01, 0x80, 0x00, 0x00, 0x04, 0x05, 0x0a, 0x81, 0x8a, 0x8b, 0x96}
var GetStripSettingsPacket = []byte{0x00, 0x02, 0x80, 0x00, 0x00, 0x05, 0x06, 0x0a, 0x63, 0x12, 0x21, 0xf0, 0x86}
var PowerOnPacket = []byte{0x00, 0x04, 0x80, 0x00, 0x00, 0x0d, 0x0e, 0x0b, 0x3b, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x00, 0x00, 0x90}
var PowerOffPacket = []byte{0x00, 0x5b, 0x80, 0x00, 0x00, 0x0d, 0x0e, 0x0b, 0x3b, 0x24, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x00, 0x00, 0x91}

type bleOptions struct {
	device   string
	duration time.Duration
}
type connectOptions struct {
	bleOptions
	addr string
}

func defaultConnectOptions() *connectOptions {
	return &connectOptions{}
}

func newConnectCmd() *cobra.Command {
	o := defaultConnectOptions()

	cmd := &cobra.Command{
		Use:          "connect [addr]",
		Short:        "Connect to device by MAC address",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		RunE:         o.run,
	}

	cmd.Flags().StringVarP(&o.device, "device", "d", "default", "implementation of ble")
	cmd.Flags().DurationVarP(&o.duration, "duration", "w", 5*time.Second, "scanning duration")

	return cmd
}

func notificationhandler(req []byte) {
	fmt.Printf("Notified: %q [ % X ]\n", string(req), req)
}

func (o *connectOptions) run(cmd *cobra.Command, args []string) error {
	err := o.parseArgs(args)
	if err != nil {
		return err
	}

	d, err := dev.NewDevice(o.device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	log.Printf("Connecting to %s...\n", o.addr)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), o.duration))
	// Dial directly, skip Connect - it's slower (as it scans) and DOESN'T EVEN WORK!... https://github.com/go-ble/ble/pull/112
	client, err := ble.Dial(ctx, ble.NewAddr(o.addr))
	if err != nil {
		return err
	}

	_, err = client.DiscoverProfile(true)
	if err != nil {
		return fmt.Errorf("can't discover profile: %s", err)
	}
	notify := client.Profile().FindCharacteristic(ble.NewCharacteristic(ble.MustParse(NotifyUUID)))
	if notify == nil {
		return fmt.Errorf("cannot find characteristic to subscribe")
	}
	write := client.Profile().FindCharacteristic(ble.NewCharacteristic(ble.MustParse(WriteUUID)))
	if write == nil {
		return fmt.Errorf("cannot find characteristic to write")
	}

	err = client.Subscribe(notify, false, notificationhandler)
	if err != nil {
		return err
	}

	err = client.WriteCharacteristic(write, InitialPacket, false)
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	err = client.WriteCharacteristic(write, GetStripSettingsPacket, false)
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	err = client.WriteCharacteristic(write, PowerOffPacket, false)
	if err != nil {
		return err
	}

	for {
		fmt.Printf("Client side RSSI: %d\n", client.ReadRSSI())
		time.Sleep(time.Second)
	}

	return nil
}

func (o *connectOptions) parseArgs(args []string) error {
	o.addr = args[0]
	return nil
}
