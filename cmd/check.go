package cmd

import (
	"fmt"

	"github.com/Portfolio-jaime/kuma-diagnoser/internal"
	"github.com/urfave/cli/v2"
)

var CheckCommand = &cli.Command{
	Name:  "check",
	Usage: "Realiza diagnóstico del control plane y dataplanes",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "allowed-clusters",
			Aliases: []string{"c"},
			Usage:   "Lista de clusters permitidos separados por coma",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"f"},
			Usage:   "Archivo de configuración YAML (opcional, sobrescribe allowed-clusters)",
		},
	},
	Action: func(c *cli.Context) error {
		fmt.Println("🔎 Validando herramientas necesarias...")
		// Corrected call to ValidateTools
		if err := internal.ValidateTools(); err != nil {
			return cli.Exit("❌ error validando herramientas: "+err.Error(), 1)
		}

		var allowedClusters []string
		if cfgFile := c.String("config"); cfgFile != "" {
			conf, err := internal.LoadConfig(cfgFile)
			if err != nil {
				return cli.Exit("❌ Error cargando config: "+err.Error(), 1)
			}
			allowedClusters = conf.AllowedClusters
		} else {
			allowedClusters = c.StringSlice("allowed-clusters")
		}

		if len(allowedClusters) == 0 {
			return cli.Exit("❌ Debe proveer al menos un cluster permitido con --allowed-clusters o --config", 1)
		}

		fmt.Println("🔎 Validando cluster conectado...")
		if err := internal.ValidateCluster(allowedClusters); err != nil {
			return cli.Exit("❌ error validando cluster: "+err.Error(), 1)
		}

		fmt.Println("🎉 ¡Todas las validaciones pasaron exitosamente!")
		fmt.Println("🔍 Ejecutando diagnóstico...")
		return internal.RunDiagnosis()
	},
}
