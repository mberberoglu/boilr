package main

import (
	"github.com/spf13/cobra"
	"github.com/tmrts/cookie/pkg/cmd"
)

func main() {
	mainCmd := &cobra.Command{
		Use: "main",
	}

	mainCmd.AddCommand(cmd.Use)
	mainCmd.AddCommand(cmd.Save)

	mainCmd.Execute()
}
