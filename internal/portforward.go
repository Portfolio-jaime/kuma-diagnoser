package internal

import (
	"fmt"
	"os/exec"
)

func RunPortForward() error {
	fmt.Println("🔌 Port forwarding a kuma-control-plane...")
	cmd := exec.Command("kubectl", "port-forward", "-n", "kuma-system", "svc/kuma-control-plane", "5681:5681")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
