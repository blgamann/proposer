package main

import (
	"log"
	"os"
	"proposer/flags"
	"proposer/service"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Usage = "L2Output Submitter"
	app.Description = "Service for generating and proposing L2 Outputs"
	app.Flags = flags.Flags
	app.Action = service.NewOutputSubmitter()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
