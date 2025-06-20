package cmd

import (
	"github.com/tu_usuario/kuma-diagnoser/internal"
	"github.com/urfave/cli/v2"
)

var ExportCommand = &cli.Command{
	Name:  "export",
	Usage: "Exporta el diagnóstico en formato json o markdown",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Required: true},
		&cli.StringFlag{Name: "output", Aliases: []string{"o"}},
	},
	Action: func(c *cli.Context) error {
		return internal.ExportDiagnosis(c.String("format"), c.String("output"))
	},
}
