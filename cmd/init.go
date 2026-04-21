package cmd

import (
	"fmt"

	"github.com/av1nDs0uza/pipeforge/internal/config"
	"github.com/av1nDs0uza/pipeforge/internal/detector"
	"github.com/av1nDs0uza/pipeforge/internal/generator"
	"github.com/av1nDs0uza/pipeforge/internal/prompts"

	"github.com/spf13/cobra"
)

var ci string
var tier string
var force bool

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {

		// validation
		if ci != "" && (ci != "github" && ci != "gitlab" && ci != "jenkins") {
			fmt.Println("Invalid CI provider")
			return
		}

		if tier != "" && (tier != "basic" && tier != "standard") {
			fmt.Println("Invalid tier")
			return
		}

		// detect project
		ctx := detector.DetectProject(".")
		fmt.Println("✔ Detected project:", ctx.Type)

		var cfg config.Config

		if ci != "" && tier != "" {
			cfg = config.Config{
				CI:   ci,
				Tier: tier,
			}
			fmt.Println("✔ Using flags:", ci, tier)

		} else {
			cfg = prompts.GetUserChoices()

			if ci != "" {
				cfg.CI = ci
			}
			if tier != "" {
				cfg.Tier = tier
			}
		}

		fmt.Println("✔ Final config:", cfg.CI, cfg.Tier)

		generator.GenerateNewProjectWithRoot(cfg, ".")
	},
}

func init() {
	initCmd.Flags().BoolVar(&force, "force", false, "Overwrite existing CI without prompt")
	initCmd.Flags().StringVar(&ci, "ci", "", "CI provider (github/gitlab/jenkins)")
	initCmd.Flags().StringVar(&tier, "tier", "", "Template tier (basic/standard)")

	rootCmd.AddCommand(initCmd)
}
