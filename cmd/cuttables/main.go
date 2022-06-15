package main

import (
	"os"

	"github.com/nerdbergev/cuttables/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := &cli.App{
		Name: "cuttables",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start the cuttables server",
				Action: func(_ *cli.Context) error {
					s := server.New()
					return s.ListenAndServe()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run the app")
	}
}
