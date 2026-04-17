package detector

import (
	"os"
)

func IsNodeProject(path string) bool {
	_, err := os.Stat(path + "/package.json")
	return err == nil
}
