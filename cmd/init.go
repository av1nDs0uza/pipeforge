package cmd

import (
	"pipeforge/internal/generator"
	"pipeforge/internal/prompts"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {

		config := prompts.GetUserChoices()

		generator.GenerateNewProject(config)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
