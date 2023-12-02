package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var repositoryFlag string
	app := &cli.App{
		Name:  "gonion",
		Usage: "Generate GO project using hexagonal architecture (onion)",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Start generate the project",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "repository",
						Aliases:     []string{"r"},
						Usage:       "module repository, without HTTP/HTTPS",
						Destination: &repositoryFlag,
					},
				},
				Action: func(context *cli.Context) error {
					Generate(repositoryFlag)
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
