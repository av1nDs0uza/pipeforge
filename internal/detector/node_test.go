package detector

import (
	"os"
	"testing"
)

func TestIsNodeProject(t *testing.T) {

	tmpDir := t.TempDir()

	// create fake package.json
	file := tmpDir + "/package.json"
	err := os.WriteFile(file, []byte("{}"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	if !IsNodeProject(tmpDir) {
		t.Errorf("expected node project detection to be true")
	}
}

func TestHasDockerfile(t *testing.T) {

	tmpDir := t.TempDir()

	file := tmpDir + "/Dockerfile"
	os.WriteFile(file, []byte("FROM alpine"), 0644)

	if !HasDockerfile(tmpDir) {
		t.Errorf("expected docker detection to be true")
	}
}
