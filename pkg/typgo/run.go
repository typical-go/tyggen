package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// RunCmd run command
	RunCmd struct {
		Precmds []string
		Action
	}
	// StdRun standard run
	StdRun struct {
		Binary string
	}
)

//
// RunCmd
//

var _ Cmd = (*RunCmd)(nil)

// Command run
func (r *RunCmd) Command(sys *BuildSys) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,

		Action: func(cliCtx *cli.Context) error {
			for _, precmd := range r.Precmds {
				if err := sys.Run(precmd, cliCtx); err != nil {
					return err
				}
			}
			return sys.Execute(r.Action, cliCtx)
		},
	}
}

//
// StdRun
//

var _ Action = (*StdRun)(nil)

// Execute standard run
func (s *StdRun) Execute(c *Context) error {
	return c.Execute(&execkit.Command{
		Name:   s.getBinary(c),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}

func (s *StdRun) getBinary(c *Context) string {
	if s.Binary == "" {
		s.Binary = fmt.Sprintf("bin/%s", c.BuildSys.Name)
	}
	return s.Binary
}
