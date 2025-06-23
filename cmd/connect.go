// based off https://github.com/go-ble/ble/blob/master/examples/basic/scanner/main.go
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/fopina/zengge-led-ctl/pkg/client"
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

func notificationhandler(req []byte) {
	fmt.Printf("Notified: %q [ % X ]\n", string(req), req)
}

func (o *connectOptions) run(cmd *cobra.Command, args []string) error {
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

	err = c.Subscribe(notificationhandler)
	if err != nil {
		return err
	}

	err = c.SendInitialPacket()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	err = c.GetStripSettings()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	err = c.PowerOff()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	return nil
}

func (o *connectOptions) parseArgs(args []string) error {
	o.addr = args[0]
	return nil
}
