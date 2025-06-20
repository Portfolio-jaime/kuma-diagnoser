package cmd

import (
	"fmt"
	"os"

	"github.com/jaimehenao8126/kuma-diagnoser/internal"
	"github.com/urfave/cli/v2"
)

var CheckCommand = &cli.Command{
	Name:  "check",
	Usage: "Realiza diagnóstico del control plane y dataplanes",
	Action: func(c *cli.Context) error {
		fmt.Println("🔎 Validando entorno...")

		// Validar herramientas necesarias
		for _, tool := range []string{"kubectl", "kumactl"} {
			if err := internal.CheckToolInstalled(tool); err != nil {
				fmt.Println("❌", err)
				os.Exit(1)
			}
		}

		// Validar conexión al cluster EKS esperado
		expectedClusterName := "my-eks-cluster" // Cambia esto al nombre real de tu cluster
		if err := internal.ValidateEksCluster(expectedClusterName); err != nil {
			fmt.Println("❌", err)
			os.Exit(1)
		}

		fmt.Println("✅ Validaciones OK, comenzando diagnóstico...")

		return internal.RunDiagnosis()
	},
}
