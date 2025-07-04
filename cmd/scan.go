// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"fmt"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
	"github.com/fopina/zengge-led-ctl/pkg/types"
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
		Use:   "scan",
		Short: "List discoverable Zengge LED devices",
		RunE:  o.run,
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

func scanHandler(a types.ZenggeAdvertisement) {
	fmt.Println(a)
}
