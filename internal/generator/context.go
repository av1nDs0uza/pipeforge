package generator

import (
	"pipeforge/internal/detector"
)

type ProjectContext struct {
	Type   string
	Docker bool
}

func DetectProject(path string) ProjectContext {

	ctx := ProjectContext{}

	if detector.IsNodeProject(path) {
		ctx.Type = "node"
	} else if detector.IsPythonProject(path) {
		ctx.Type = "python"
	} else {
		ctx.Type = "unknown"
	}

	ctx.Docker = detector.HasDockerfile(path)

	return ctx
}
