package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/CzarSimon/mgen/pkg"
	"github.com/CzarSimon/mgen/pkg/generator"
	"github.com/urfave/cli"
)

const (
	appName        = "mgen"
	appUsage       = "Code gerator for data models"
	appDescription = "mgen is a tool for generating data models in multiple languages and serialization formats."
	appVersion     = "0.1"
)

var (
	errNoFilepath = errors.New("No filepath specified")
)

func main() {
	err := setupApp().Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Description = appDescription
	app.Version = appVersion
	app.Action = generateModel
	return app
}

func generateModel(c *cli.Context) error {
	filepath := getFilepathArg(c)
	schema, err := pkg.ReadSchema(filepath)
	checkErr(err)
	gen := generator.NewGo()
	block, err := gen.Generate(schema)
	checkErr(err)
	fmt.Printf("%s\n", block)
	return nil
}

func getFilepathArg(c *cli.Context) string {
	filepath := c.Args().Get(0)
	if filepath == "" {
		checkErr(errNoFilepath)
	}
	return filepath
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
