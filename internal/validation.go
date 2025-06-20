package internal

import (
	"fmt"
	"os/exec"
	"strings"

	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AllowedClusters []string `yaml:"allowedClusters"`
}

func ValidateTools() error {
	fmt.Println("🔎 Iniciando validación de herramientas necesarias...")
	// Ejemplo: Validar si kubectl está instalado
	_, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("❌ kubectl no está instalado.")
		return fmt.Errorf("kubectl no está instalado: %v", err)
	}
	fmt.Println("✅ ¡Herramienta kubectl validada exitosamente! Está lista para usar.")
	return nil
}

func ValidateCluster(allowedClusters []string) error {
	fmt.Println("🔎 Iniciando validación del cluster conectado...")
	out, err := exec.Command("kubectl", "config", "current-context").Output()
	if err != nil {
		fmt.Println("❌ Error obteniendo contexto kubectl.")
		return fmt.Errorf("error obteniendo contexto kubectl: %v", err)
	}
	current := strings.TrimSpace(string(out))

	// Extraer el nombre del cluster del ARN si es necesario
	currentParts := strings.Split(current, "/")
	clusterName := currentParts[len(currentParts)-1]

	for _, c := range allowedClusters {
		if clusterName == c {
			fmt.Printf("✅ ¡Cluster '%s' validado exitosamente! Está dentro de los clusters permitidos.\n", clusterName)
			return nil
		}
	}
	fmt.Printf("❌ Contexto kubectl actual '%s' no está entre clusters permitidos: %v\n", clusterName, allowedClusters)
	return fmt.Errorf("contexto kubectl actual '%s' no está entre clusters permitidos: %v", clusterName, allowedClusters)
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
