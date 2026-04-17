package cmd

import (
	"github.com/av1nDs0uza/pipeforge/internal/generator"
	"github.com/av1nDs0uza/pipeforge/internal/prompts"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {

		config := prompts.GetUserChoices()

		generator.InjectIntoExisting(config, ".")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
