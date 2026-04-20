package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/av1nDs0uza/pipeforge/internal/config"
	"github.com/av1nDs0uza/pipeforge/internal/detector"
	"github.com/av1nDs0uza/pipeforge/internal/templates"
)

func GenerateNewProjectWithRoot(cfg config.Config, root string) {

	fmt.Println("Creating new CI/CD setup...")

	ctx := detector.DetectProject(root)

	fileName := fmt.Sprintf("%s-%s-%s.yml", cfg.CI, ctx.Type, cfg.Tier)

	// READ template
	data, err := templates.Files.ReadFile(fileName)
	if err != nil {
		fmt.Println("Template not found:", err)
		return
	}

	// correct target path
	target := filepath.Join(root, ".github", "workflows", "ci.yml")

	// create folders
	err = os.MkdirAll(filepath.Dir(target), 0755)
	if err != nil {
		fmt.Println("Failed to create folders:", err)
		return
	}

	// write file
	err = os.WriteFile(target, data, 0644)
	if err != nil {
		fmt.Println("Write failed:", err)
		return
	}

	fmt.Println("✔ CI/CD pipeline created at .github/workflows/ci.yml")
}

func InjectIntoExisting(cfg config.Config, root string) {

	fmt.Println("Injecting CI/CD into existing project...")

	ctx := detector.DetectProject(root)

	fileName := fmt.Sprintf("%s-%s-%s.yml", cfg.CI, ctx.Type, cfg.Tier)

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
