package cmd

import (
	"log"
	"strconv"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
	"github.com/spf13/cobra"
)

type colorOptions struct {
	connectOptions
	red,
	green,
	blue byte
}

func newColorCmd() *cobra.Command {
	o := &colorOptions{}

	cmd := &cobra.Command{
		Use:   "color [addr] [red] [green] [blue]",
		Short: "Set strip color by MAC address, using RGB (0-255)",
		Args:  cobra.ExactArgs(4),
		RunE:  o.run,
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

	return c.SetRGBBytes(o.red, o.green, o.blue)
}

func (o *colorOptions) parseArgs(args []string) error {
	err := o.connectOptions.parseArgs(args)
	if err != nil {
		return err
	}
	fields := []*byte{&o.red, &o.green, &o.blue}
	for i := 0; i < 3; i++ {
		val, err := strconv.ParseUint(args[i+1], 10, 8)
		if err != nil {
			return err
		}
		*fields[i] = byte(val)
	}
	return nil
}
