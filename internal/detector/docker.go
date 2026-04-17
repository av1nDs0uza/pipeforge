package detector

import "os"

func HasDockerfile(path string) bool {
	_, err := os.Stat(path + "/Dockerfile")
	return err == nil
}
