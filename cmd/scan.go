// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"fmt"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
	"github.com/go-ble/ble"
	"github.com/spf13/cobra"
)

type scanOptions struct {
	bleOptions
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
	c, err := client.NewZenggeClient(o.device)
	if err != nil {
		return err
	}
	// Scan for specified durantion, or until interrupted by user.
	fmt.Printf("Scanning for %s...\n", o.duration)
	return c.Scan(o.duration, o.duplicates, scanHandler)
}

func scanHandler(a ble.Advertisement) {
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
