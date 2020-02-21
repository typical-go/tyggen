package typicalgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// TypicalGo is app of typical-go
type TypicalGo struct{}

// New of Typical-Go
func New() *TypicalGo {
	return &TypicalGo{}
}

// Run the typical-go
func (t *TypicalGo) Run(d *typcore.Descriptor) (err error) {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version

	app.Commands = []*cli.Command{
		{
			Name:      "new",
			Usage:     "Construct New Project",
			UsageText: "app new [Package]",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "blank", Usage: "Create blank new project"},
			},
			Action: func(c *cli.Context) (err error) {
				pkg := c.Args().First()
				if pkg == "" {
					return cli.ShowCommandHelp(c, "new")
				}
				return t.constructProject(c.Context, pkg)
			},
		},
	}
	return app.Run(os.Args)
}
