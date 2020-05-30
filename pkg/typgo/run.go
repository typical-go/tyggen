package typgo

import (
	"errors"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// Runner responsible to run
	Runner interface {
		Run(*Context) error
	}

	// StdRun standard run
	StdRun struct{}
)

var _ Runner = (*StdRun)(nil)

//
// StdRun
//

// Run for standard typical project
func (*StdRun) Run(c *Context) error {
	return execute(c, &execkit.Command{
		Name:   typvar.AppBin(c.Descriptor.Name),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
}

//
// command
//

func cmdRun(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFn("RUN", run),
	}
}

func run(c *Context) error {
	if c.Run == nil {
		return errors.New("run is missing")
	}

	if err := compile(c); err != nil {
		return err
	}

	return c.Run.Run(c)
}
