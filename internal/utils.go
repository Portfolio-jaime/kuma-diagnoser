package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunKubectl(cmd string) []string {
	return run("kubectl", cmd)
}

func RunKumactl(cmd string) []string {
	return run("kumactl", cmd)
}

func run(tool, cmd string) []string {
	c := exec.Command("bash", "-c", fmt.Sprintf("%s %s", tool, cmd))
	out, err := c.Output()
	if err != nil {
		return []string{}
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}
