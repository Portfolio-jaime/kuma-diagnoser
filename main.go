package main

import (
	"log"
	"os"

	"github.com/jaimehenao8126/kuma-diagnoser/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "kuma-diagnoser",
		Usage: "Diagnósticos para Kuma Service Mesh",
		Commands: []*cli.Command{
			cmd.CheckCommand,
			cmd.ExportCommand,
			cmd.PortForwardCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
