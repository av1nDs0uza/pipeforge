package detector

import "os"

func IsGoProject(path string) bool {
	_, err := os.Stat(path + "/go.mod")
	return err == nil
}
