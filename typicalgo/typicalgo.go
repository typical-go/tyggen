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
			Name: "wrap-me",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "tmp", Required: true},
				&cli.StringFlag{Name: "project-package"},
			},
			Action: func(c *cli.Context) (err error) {
				var (
					root           string
					tmp            = c.String("tmp")
					projectPackage = c.String("project-package")
				)

				if root, err = os.Getwd(); err != nil {
					return err
				}

				if projectPackage == "" {
					projectPackage = typcore.RetrieveProjectPackage(root)
				}

				return wrapMe(&wrapContext{
					Context:        typcore.CreateContext(d),
					Cli:            c,
					tmp:            tmp,
					projectPackage: projectPackage,
				})
			},
		},
	}
	return app.Run(os.Args)
}
