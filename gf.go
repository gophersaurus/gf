package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
)

// USAGE
//
// gf [global options..] command [command options...]          [arguments...]
//    -v, -verbose       install -g="git@github.git"           project
//                               -git="https://github.com.git"
//

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// start a new cli application.
	app := cli.NewApp()

	// define the cli application metadata.
	app.Name = "gf"
	app.Usage = "A CLI tool for starting Gophersaurus projects."
	app.Version = "0.0.1"
	app.Author = "Jack Spirou"
	app.Email = "jack.spirou@me.com"

	// define the list of global options.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose Mode: Print stuff as it happens.",
		},
	}

	// define the list of commands.
	app.Commands = []cli.Command{
		{
			Name:        "new",
			Usage:       "Create a new Gophersaurus project.",
			ShortName:   "n",
			Description: "Use this command to create a new Gophersaurus project.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "g, git",
					Usage: "Git: Create the project with a git repository and create upstream.",
				},
			},
			Action: func(c *cli.Context) {
				if err := create(c.Args().First(), c.String("g"), c.GlobalBool("verbose")); err != nil {
					panic(err)
				}
			},
		},
	}

	// define the default action.
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	// execute the cli command given.
	app.Run(os.Args)
}
