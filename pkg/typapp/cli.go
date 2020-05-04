package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

func createAppCli(a *App, d *typcore.Descriptor) *cli.App {
	c := &Context{
		Descriptor: d,
		App:        a,
		Logger: typlog.Logger{
			Name: d.Name,
		},
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Before = func(*cli.Context) (err error) {
		if configFile := os.Getenv("CONFIG"); configFile != "" {
			_, err = typcfg.Load(configFile)
		}
		return
	}
	app.Version = d.Version
	app.Action = c.ActionFunc(a.main)
	app.Commands = a.Commands(c)
	return app
}