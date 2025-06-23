// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"log"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
	"github.com/spf13/cobra"
)

type colorOptions struct {
	connectOptions
	color string
}

func newColorCmd() *cobra.Command {
	o := &colorOptions{}

	cmd := &cobra.Command{
		Use:          "power [addr] [state]",
		Short:        "Power device by MAC address, 1 for ON and 0 for OFF",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(2),
		RunE:         o.run,
	}

	cmd.Flags().StringVarP(&o.device, "device", "d", "default", "implementation of ble")
	cmd.Flags().DurationVarP(&o.duration, "duration", "w", 5*time.Second, "scanning duration")

	return cmd
}

func (o *colorOptions) run(cmd *cobra.Command, args []string) error {
	err := o.parseArgs(args)
	if err != nil {
		return err
	}

	c, err := client.NewZenggeClient(o.device)
	if err != nil {
		return err
	}

	log.Printf("Connecting to %s...\n", o.addr)
	err = c.Connect(o.addr, o.duration)
	if err != nil {
		return err
	}
	return nil
}

func (o *colorOptions) parseArgs(args []string) error {
	err := o.connectOptions.parseArgs(args)
	if err != nil {
		return err
	}
	return nil
}
