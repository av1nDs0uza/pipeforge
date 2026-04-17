package detector

import "os"

func IsPythonProject(path string) bool {
	_, err := os.Stat(path + "/requirements.txt")
	return err == nil
}
