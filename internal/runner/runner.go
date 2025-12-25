package runner

import (
	"os/exec"
)

func RunCommand(cmd []string) (stdout string, stderr string, exitCode int) {
	c := exec.Command(cmd[0], cmd[1:]...)
	outBytes, err := c.CombinedOutput()
	stdout = string(outBytes)
	stderr = string(outBytes)

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = 1
		}
	} else {
		exitCode = 0
	}

	return
}
