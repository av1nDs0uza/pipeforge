package cmd

import (
	"fmt"

	"pipeforge/internal/generator"
	"pipeforge/internal/prompts"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {

		// detect current folder
		ctx := generator.DetectProject(".")

		// user input (Vite-style override)
		config := prompts.GetUserChoices()

		fmt.Println("Detected project:", ctx.Type)

		generator.GenerateNewProjectWithRoot(config, ".")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
