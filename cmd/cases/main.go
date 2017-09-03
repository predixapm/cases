package main

import (
	"fmt"
	"os"

	"github.com/teshguru/cases"
	"github.com/caarlos0/spin"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "cases"
	app.Version = version
	app.Author = "Carlos Alexandro Becker (root@carlosbecker.com)"
	app.Usage = "This is an cases app written in Go"
	app.Action = func(c *cli.Context) error {
		spin := spin.New("\033[36m %s Working...\033[m")
		spin.Start()
		err := cases.Foo()
		spin.Stop()
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	}
	_ = app.Run(os.Args)
}
