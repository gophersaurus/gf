package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
)

// USAGE
//
// gf [global options..] command [command options...]          [arguments...]
//    -v, -verbose       new     -origin="git@github.git"      project-name
//                       n       -o="https://github.com.git"
// 															 -s -skip
//															 -sg -skipgit
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
					Name:  "o, origin",
					Usage: "Origin: Create the project with a git remote origin.",
				},
				cli.BoolFlag{
					Name:  "s, skip",
					Usage: "Skip: Skip any warnings that would require user input.",
				},
				cli.BoolFlag{
					Name:  "sg, skipgit",
					Usage: "Skip Git: Skip default git settings.",
				},
			},
			Action: func(c *cli.Context) {
				if err := create(c.Args().First(), c.String("o"), c.GlobalBool("verbose"), c.Bool("skip"), c.Bool("skipgit")); err != nil {
					fmt.Println(err.Error())
					panic(err)
				}
			},
		},
		{
			Name:        "update",
			Usage:       "Update the Gophersaurus framework.",
			ShortName:   "u",
			Description: "Use this command to update the Gophersaurus framework in the current directory.",
			Action: func(c *cli.Context) {
				if err := update(); err != nil {
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
