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

type connectOptions struct {
	device     string
	duration   time.Duration
	duplicates bool
}

func defaultConnectOptions() *connectOptions {
	return &connectOptions{}
}

func newConnectCmd() *cobra.Command {
	o := defaultConnectOptions()

	cmd := &cobra.Command{
		Use:          "connect",
		Short:        "List discoverable Zengge LED devices",
		SilenceUsage: true,
		RunE:         o.run,
	}

	cmd.Flags().StringVarP(&o.device, "device", "d", "default", "implementation of ble")
	cmd.Flags().DurationVarP(&o.duration, "duration", "w", 5*time.Second, "scanning duration")
	cmd.Flags().BoolVarP(&o.duplicates, "dup", "", true, "allow duplicate reported")

	return cmd
}

func (o *connectOptions) run(cmd *cobra.Command, args []string) error {
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
