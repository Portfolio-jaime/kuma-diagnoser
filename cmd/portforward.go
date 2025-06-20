package cmd

import (
	"github.com/tu_usuario/kuma-diagnoser/internal"
	"github.com/urfave/cli/v2"
)

var PortForwardCommand = &cli.Command{
	Name:  "port-forward",
	Usage: "Ejecuta port-forward al control plane de Kuma",
	Action: func(c *cli.Context) error {
		return internal.RunPortForward()
	},
}
