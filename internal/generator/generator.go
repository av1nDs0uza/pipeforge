package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/av1nDs0uza/pipeforge/internal/config"
	"github.com/av1nDs0uza/pipeforge/internal/templates"
)

func GenerateNewProjectWithRoot(cfg config.Config, root string) {

	fmt.Println("Creating new CI/CD setup...")

	fileName := fmt.Sprintf("%s-%s.yml", cfg.CI, cfg.Tier)

	// READ from embedded templates
	data, err := templates.Files.ReadFile(fileName)
	if err != nil {
		fmt.Println("Template not found:", err)
		return
	}

	// WRITE to root directory
	target := filepath.Join(root, ".ci.yml")

	err = os.WriteFile(target, data, 0644)
	if err != nil {
		fmt.Println("Write failed:", err)
		return
	}

	fmt.Println("✔ CI/CD pipeline created")
}

func InjectIntoExisting(cfg config.Config, root string) {

	fmt.Println("Injecting CI/CD into existing project...")

	fileName := fmt.Sprintf("%s-%s.yml", cfg.CI, cfg.Tier)

	data, err := templates.Files.ReadFile(fileName)
	if err != nil {
		fmt.Println("Template not found:", err)
		return
	}

	target := filepath.Join(root, ".github", "workflows", "ci.yml")

	err = os.MkdirAll(filepath.Dir(target), 0755)
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
