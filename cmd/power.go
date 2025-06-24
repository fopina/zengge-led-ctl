// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
	"github.com/spf13/cobra"
)

type powerOptions struct {
	connectOptions
	state bool
}

func newPowerCmd() *cobra.Command {
	o := &powerOptions{}

	cmd := &cobra.Command{
		Use:   "power [addr] [state]",
		Short: "Power device by MAC address, 1 for ON and 0 for OFF",
		Args:  cobra.ExactArgs(2),
		RunE:  o.run,
	}

	cmd.Flags().StringVarP(&o.device, "device", "d", "default", "implementation of ble")
	cmd.Flags().DurationVarP(&o.duration, "duration", "w", 5*time.Second, "scanning duration")

	return cmd
}

func (o *powerOptions) run(cmd *cobra.Command, args []string) error {
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

	if o.state {
		return c.PowerOn()
	}
	return c.PowerOff()
}

func (o *powerOptions) parseArgs(args []string) error {
	err := o.connectOptions.parseArgs(args)
	if err != nil {
		return err
	}
	switch strings.ToLower(args[1]) {
	case "on", "1":
		o.state = true
	case "off", "0":
		o.state = false
	default:
		return fmt.Errorf("invalid state: %s", args[1])
	}
	return nil
}
