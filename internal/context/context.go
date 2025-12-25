package context

import (
	"os"
)

type Context struct {
	Command  []string
	Cwd      string
	Env      map[string]string
	Stdout   string
	Stderr   string
	ExitCode int
}

func BuildContext(cmd []string) *Context {
	cwd, _ := os.Getwd()
	envMap := make(map[string]string)
	for _, e := range os.Environ() {
		pair := splitOnce(e, "=")
		envMap[pair[0]] = pair[1]
	}

	return &Context{
		Command: cmd,
		Cwd:     cwd,
		Env:     envMap,
	}
}

func splitOnce(s, sep string) [2]string {
	i := 0
	for i < len(s) && s[i] != sep[0] {
		i++
	}
	if i == len(s) {
		return [2]string{s, ""}
	}
	return [2]string{s[:i], s[i+1:]}
}
