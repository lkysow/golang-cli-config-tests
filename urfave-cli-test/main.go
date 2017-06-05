package main

import (
	"os"

    "gopkg.in/urfave/cli.v1"
	"fmt"
	"github.com/urfave/cli/altsrc"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "server test",
			Action: func(c *cli.Context) error {
				conf := make(map[string]string)
				conf["config"] = c.String("config")
				conf["gh-username"] = c.GlobalString("gh-username")
				fmt.Printf("config: %v", conf)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "gh-username"},
				altsrc.NewStringFlag(cli.StringFlag{Name: "gh-username"}),
			},
		},
	}
	flags := []cli.Flag{
		cli.StringFlag{Name: "config", Usage: "config file"},
	}

    app.Before = func(c *cli.Context) error {
		if c.String("config") != "" {
			return altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))(c)
		}
		return nil
	}
	app.Flags = flags
	app.Run(os.Args)
}
