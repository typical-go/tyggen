package typgo

import (
	"errors"
	"os"
	"regexp"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"go.uber.org/dig"
)

type (

	// Descriptor describe the project
	Descriptor struct {
		// Name of the project (OPTIONAL). It should be a characters with/without underscore or dash.
		// By default, project name is same with project folder
		Name string
		// Description of the project (OPTIONAL).
		Description string
		// Version of the project (OPTIONAL). By default it is 0.0.1
		Version string

		EntryPoint interface{}
		Layouts    []string

		Test    Tester
		Compile Compiler
		Run     Runner
		Release Releaser
		Utility Utility

		SkipPrecond bool
		Configurer  Configurer
	}
)

var _ typcore.AppLauncher = (*Descriptor)(nil)
var _ typcore.BuildLauncher = (*Descriptor)(nil)

// LaunchApp to launch the app
func (d *Descriptor) LaunchApp() (err error) {
	if err = d.Validate(); err != nil {
		return
	}
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		_, err = LoadConfig(configFile)
	}

	di := dig.New()
	if err = setDependencies(di, d); err != nil {
		return
	}

	errs := common.GracefulRun(start(di, d), stop(di))
	return errs.Unwrap()
}

// LaunchBuild to launch the build tool
func (d *Descriptor) LaunchBuild() (err error) {
	return launchBuild(d)
}

// Validate context
func (d *Descriptor) Validate() (err error) {
	if d.Version == "" {
		d.Version = "0.0.1"
	}

	if !ValidateName(d.Name) {
		return errors.New("Descriptor: bad name")
	}

	return
}

// ValidateName to validate valid descriptor name
func ValidateName(name string) bool {
	if name == "" {
		return false
	}

	r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
	if !r.MatchString(name) {
		return false
	}
	return true
}
