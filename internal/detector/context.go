package detector

type ProjectContext struct {
	Type   string
	Docker bool
}

func DetectProject(path string) ProjectContext {

	ctx := ProjectContext{}

	if IsNodeProject(path) {
		ctx.Type = "node"
	} else if IsPythonProject(path) {
		ctx.Type = "python"
	} else if IsGoProject(path) {
		ctx.Type = "go"
	} else {
		ctx.Type = "node" // safe fallback
	}

	ctx.Docker = HasDockerfile(path)

	return ctx
}
