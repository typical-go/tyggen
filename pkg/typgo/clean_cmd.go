package typgo

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// CleanCmd command clean
	CleanCmd struct {
		Action
	}
	// StdClean standard clean
	StdClean struct {
		Paths []string
	}
)

//
// CleanCmd
//

var _ Cmd = (*CleanCmd)(nil)

// Command clean
func (c *CleanCmd) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:   "clean",
		Usage:  "Clean the project",
		Action: b.ActionFn(c),
	}
}

//
// StdClean
//

var _ Action = (*StdClean)(nil)

// Execute standard clean
func (s *StdClean) Execute(c *Context) error {
	for _, path := range s.GetPaths() {
		if err := os.RemoveAll(path); err != nil {
			fmt.Printf("Failed removing %s\n", path)
		} else {
			fmt.Printf("Removing %s\n", path)
		}

	}
	// removeAll(c, TypicalTmp)
	return nil
}

// GetPaths return paths to be clean
func (s *StdClean) GetPaths() []string {
	if len(s.Paths) < 1 {
		s.Paths = []string{TypicalTmp}
	}
	return s.Paths
}