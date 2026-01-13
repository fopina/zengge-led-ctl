package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
	"github.com/spf13/cobra"
)

type colorNameOptions struct {
	connectOptions
	name string
}

func newColorCmd() *cobra.Command {
	o := &colorNameOptions{}

	cmd := &cobra.Command{
		Use:   "color [addr] [name]",
		Short: "Set strip color by MAC address, using a color name (e.g. red, green)",
		Args:  cobra.ExactArgs(2),
		RunE:  o.run,
	}

	cmd.Flags().StringVarP(&o.device, "device", "d", "default", "implementation of ble")
	cmd.Flags().DurationVarP(&o.duration, "duration", "w", 5*time.Second, "scanning duration")

	return cmd
}

func (o *colorNameOptions) run(cmd *cobra.Command, args []string) error {
	if err := o.parseArgs(args); err != nil {
		return err
	}

	r, g, b, ok := nameToRGB(o.name)
	if !ok {
		return fmt.Errorf("unknown color name: %s", o.name)
	}

	c, err := client.NewZenggeClient(o.device)
	if err != nil {
		return err
	}

	log.Printf("Connecting to %s...\n", o.addr)
	if err := c.Connect(o.addr, o.duration); err != nil {
		return err
	}

	return c.SetRGBBytes(r, g, b)
}

func (o *colorNameOptions) parseArgs(args []string) error {
	if err := o.connectOptions.parseArgs(args); err != nil {
		return err
	}
	o.name = strings.ToLower(args[1])
	return nil
}

// nameToRGB provides a small set of common color name mappings.
// Extend as needed.
func nameToRGB(name string) (byte, byte, byte, bool) {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "red":
		return 255, 0, 0, true
	case "green":
		return 0, 255, 0, true
	case "blue":
		return 0, 0, 255, true
	case "white":
		return 255, 255, 255, true
	case "black", "off":
		return 0, 0, 0, true
	case "yellow":
		return 255, 255, 0, true
	case "cyan", "aqua":
		return 0, 255, 255, true
	case "magenta", "fuchsia":
		return 255, 0, 255, true
	case "orange":
		return 255, 165, 0, true
	case "purple":
		return 128, 0, 128, true
	case "pink":
		return 255, 192, 203, true
	case "lime":
		return 191, 255, 0, true
	case "teal":
		return 0, 128, 128, true
	case "navy":
		return 0, 0, 128, true
	case "gray", "grey":
		return 128, 128, 128, true
	default:
		return 0, 0, 0, false
	}
}
