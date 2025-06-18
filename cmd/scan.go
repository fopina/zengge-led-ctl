package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type scanOptions struct {
	multiply bool
	add      bool
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

	cmd.Flags().BoolVarP(&o.multiply, "multiply", "m", o.multiply, "multiply")
	cmd.Flags().BoolVarP(&o.add, "add", "a", o.add, "add")

	return cmd
}

func (o *scanOptions) run(cmd *cobra.Command, args []string) error {
	fmt.Println("oi")

	return nil
}
