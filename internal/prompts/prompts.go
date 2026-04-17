package prompts

import (
	"github.com/av1nDs0uza/pipeforge/internal/config"

	"github.com/manifoldco/promptui"
)

func GetUserChoices() config.Config {

	ciPrompt := promptui.Select{
		Label: "Select CI Tool",
		Items: []string{"github", "gitlab", "jenkins"},
	}

	_, ci, _ := ciPrompt.Run()

	tierPrompt := promptui.Select{
		Label: "Select Tier",
		Items: []string{"basic", "standard", "premium"},
	}

	_, tier, _ := tierPrompt.Run()

	dockerPrompt := promptui.Select{
		Label: "Use Docker?",
		Items: []string{"yes", "no"},
	}

	_, docker, _ := dockerPrompt.Run()

	return config.Config{
		CI:     ci,
		Tier:   tier,
		Docker: docker == "yes",
	}
}
