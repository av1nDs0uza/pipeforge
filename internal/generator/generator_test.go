package generator

import (
	"os"
	"path/filepath"
	"testing"

	"pipeforge/internal/config"
)

func TestGenerateNewProject(t *testing.T) {

	tmpDir := t.TempDir()

	cfg := config.Config{
		CI:   "github",
		Tier: "basic",
	}

	GenerateNewProjectWithRoot(cfg, tmpDir)

	target := filepath.Join(tmpDir, ".ci.yml")

	if _, err := os.Stat(target); err != nil {
		t.Errorf("expected .ci.yml to be created")
		t.Errorf("expected file at %s", target)
	}
}
