package utils

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/samber/lo"
)

func EvalSh(code string, env map[string]string) (string, string, int, error) {
	cmd := exec.Command("sh", "-c", code)
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env,
		lo.MapToSlice(env, func(k string, v string) string {
			return k + "=" + env[k]
		})...,
	)

	// new string writer
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stdout.String(), stderr.String(), cmd.ProcessState.ExitCode(), err
	// return stdout.String(), stderr.String(), err
}
