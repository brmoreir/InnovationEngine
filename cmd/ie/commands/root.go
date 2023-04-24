package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "ie",
	Short: "The innovation engine.",
}

// Entrypoint into the Innovation Engine CLI.
func ExecuteCLI() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}