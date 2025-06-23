package main

import (
	"fmt"
	"os"

	"github.com/fopina/zengge-led-ctl/cmd"
)

var version = "dev"

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
