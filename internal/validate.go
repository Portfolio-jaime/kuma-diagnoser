package internal

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func CheckToolInstalled(tool string) error {
	_, err := exec.LookPath(tool)
	if err != nil {
		return fmt.Errorf("herramienta requerida no encontrada: %s. Por favor instálala y asegúrate que esté en tu PATH", tool)
	}
	return nil
}

func GetCurrentKubectlContext() (string, error) {
	out, err := exec.Command("kubectl", "config", "current-context").Output()
	if err != nil {
		return "", fmt.Errorf("falló obtener el contexto actual de kubectl: %w", err)
	}
	context := strings.TrimSpace(string(out))
	if context == "" {
		return "", errors.New("no se detectó ningún contexto kubectl configurado")
	}
	return context, nil
}

func ValidateEksCluster(expectedClusterName string) error {
	context, err := GetCurrentKubectlContext()
	if err != nil {
		return err
	}

	if !strings.Contains(context, expectedClusterName) {
		return fmt.Errorf("conectado al contexto '%s', pero se esperaba conexión al cluster '%s'", context, expectedClusterName)
	}

	// También podemos probar un comando simple para confirmar conectividad
	out, err := exec.Command("kubectl", "get", "nodes").Output()
	if err != nil {
		return fmt.Errorf("error ejecutando 'kubectl get nodes': %w", err)
	}

	if len(out) == 0 {
		return errors.New("no se encontraron nodos en el cluster. Posible problema de conexión o permisos")
	}

	return nil
}
