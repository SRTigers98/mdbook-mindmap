package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// supportsCmd represents the supports command
var supportsCmd = &cobra.Command{
	Use:   "supports",
	Short: "Checks if the given renderer is supported",
	Long: `Returns with exit code 0 if the given renderer is supported.
Otherwise the command returns with an exit code of 1.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "html" {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(supportsCmd)
}
