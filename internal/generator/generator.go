package generator

import (
	"fmt"
	"os"

	"pipeforge/internal/config"
)

func GenerateNewProject(cfg config.Config) {

	fmt.Println("Creating new CI/CD setup...")

	file := fmt.Sprintf("templates/%s-%s.yml", cfg.CI, cfg.Tier)

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Template not found:", err)
		return
	}

	os.WriteFile(".ci.yml", data, 0644)

	fmt.Println("✔ CI/CD pipeline created")
}

func InjectIntoExisting(config config.Config) {

	fmt.Println("Injecting CI/CD into existing project...")

	file := fmt.Sprintf("templates/%s-%s.yml", config.CI, config.Tier)

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Template not found:", err)
		return
	}

	// IMPORTANT: real GitHub path
	target := ".github/workflows/ci.yml"

	err = os.MkdirAll(".github/workflows", 0755)
	if err != nil {
		fmt.Println("Failed to create folders:", err)
		return
	}

	err = os.WriteFile(target, data, 0644)
	if err != nil {
		fmt.Println("Write failed:", err)
		return
	}

	fmt.Println("✔ CI/CD injected into existing project")
}
