package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fiat"
	app.Usage = "free input analysis tool"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "capture",
			ShortName: "c",
			Usage:     "capture input",
			Action: func(c *cli.Context) {
				fmt.Println("TODO")
			},
		},
	}
	app.Action = func(c *cli.Context) {
		fmt.Println("TODO")
	}
	app.Run(os.Args)
}
