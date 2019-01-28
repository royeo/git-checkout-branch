package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func cmdOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprint(stderr(), string(b))
		os.Exit(1)
	}
	return string(b)
}

func cmdRun(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = stdout()
	cmd.Stderr = stderr()
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

func stdout() io.Writer {
	return os.Stdout
}

func stderr() io.Writer {
	return os.Stderr
}
