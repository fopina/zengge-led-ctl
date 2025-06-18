// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/spf13/cobra"
)

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
