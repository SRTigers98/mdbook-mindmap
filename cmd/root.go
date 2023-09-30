package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/srtigers98/mdbook-mindmap/service"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mdbook-mindmap",
	Short: "A mdbook preprocessor to create mindmaps from the book structure",
	Long:  "A mdbook preprocessor to create mindmaps from the book structure.",
	Run: func(cmd *cobra.Command, args []string) {
		var input []map[string]any
		err := json.NewDecoder(os.Stdin).Decode(&input)
		cobra.CheckErr(err)

		book := service.ProcessBook(input[0], input[1])

		output, err := json.Marshal(book)
		cobra.CheckErr(err)

		fmt.Println(string(output))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
