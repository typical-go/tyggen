package typgo

import "github.com/urfave/cli/v2"

func cmdRun(c *BuildTool) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFunc("RUN", run),
	}
}

func run(c *Context) (err error) {
	for _, module := range c.BuildTool.BuildSequences {
		if runner, ok := module.(Runner); ok {
			if err = runner.Run(c); err != nil {
				return
			}
		}
	}
	return
}
