// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type scanOptions struct {
	device     string
	duration   time.Duration
	duplicates bool
}

func defaultScanOptions() *scanOptions {
	return &scanOptions{}
}

func newScanCmd() *cobra.Command {
	o := defaultScanOptions()

	cmd := &cobra.Command{
		Use:          "scan",
		Short:        "List discoverable Zengge LED devices",
		SilenceUsage: true,
		RunE:         o.run,
	}

	cmd.Flags().StringVarP(&o.device, "device", "d", "default", "implementation of ble")
	cmd.Flags().DurationVarP(&o.duration, "duration", "w", 5*time.Second, "scanning duration")
	cmd.Flags().BoolVarP(&o.duplicates, "dup", "", true, "allow duplicate reported")

	return cmd
}

func (o *scanOptions) run(cmd *cobra.Command, args []string) error {
	d, err := dev.NewDevice(o.device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	// Scan for specified durantion, or until interrupted by user.
	fmt.Printf("Scanning for %s...\n", o.duration)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), o.duration))
	chkErr(ble.Scan(ctx, o.duplicates, advHandler, nil))

	return nil
}

func advHandler(a ble.Advertisement) {
	if !strings.HasPrefix(a.LocalName(), "LEDnetWF") {
		return
	}

	if a.Connectable() {
		fmt.Printf("[%s] C %3d:", a.Addr(), a.RSSI())
	} else {
		fmt.Printf("[%s] N %3d:", a.Addr(), a.RSSI())
	}
	comma := ""
	if len(a.LocalName()) > 0 {
		fmt.Printf(" Name: %s", a.LocalName())
		comma = ","
	}
	if len(a.Services()) > 0 {
		fmt.Printf("%s Svcs: %v", comma, a.Services())
		comma = ","
	}
	if len(a.ManufacturerData()) > 0 {
		fmt.Printf("%s MD: %X", comma, a.ManufacturerData())
	}
	fmt.Printf("\n")
}

func chkErr(err error) {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}
