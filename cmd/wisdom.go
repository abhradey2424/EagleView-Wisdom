package main

import (
	"fmt"
	"os"

	"github.com/abhradey2424/EagleView-Wisdom/lib/wisdom"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "wisdom"
	app.Usage = "Dispense some programming fortune cookies"
	app.UsageText = "wisdom dispense"
	app.Version = wisdom.Version

	app.Commands = []cli.Command{
		cli.Command{
			Name:      "dispense",
			Usage:     "Dispense some programming fortune cookies",
			UsageText: "wisdom dispense",
			Action:    DispenseWisdom,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
}

func DispenseWisdom(context *cli.Context) error {
	fmt.Println("Dispense some wisdom")
	return nil
}
