package app_test

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestSetup(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	os.Mkdir("some-pkg", 0777)
	defer os.RemoveAll("some-pkg")

	err := app.Setup(cliContext([]string{
		"-project-pkg=some-pkg",
	}))
	require.NoError(t, err)

	b, _ := ioutil.ReadFile("some-pkg/typicalw")
	require.Equal(t, `#!/bin/bash

set -e

TYPTMP=.typical-tmp
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	echo "Build typical-go"
	go build -o $TYPGO github.com/typical-go/typical-go
fi

$TYPGO run \
	-project-pkg="some-pkg" \
	-typical-build="tools/typical-build" \
	-typical-tmp=$TYPTMP \
	$@
`, string(b))

	require.Equal(t, "Create 'some-pkg/typicalw'\n", output.String())
}

func TestSetup_GetParamError(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go list -m", ReturnError: errors.New("some-error")},
	})
	defer unpatch(t)

	os.Mkdir(".typical-tmp", 0777)
	defer os.RemoveAll(".typical-tmp")

	err := app.Setup(cliContext([]string{}))
	require.EqualError(t, err, "some-error: ")
}

func TestSetup_WithGomodFlag(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go mod init somepkg"},
	})
	defer unpatch(t)

	os.Mkdir("somepkg", 0777)
	defer os.RemoveAll("somepkg")

	err := app.Setup(cliContext([]string{
		"-project-pkg=somepkg",
		"-go-mod",
	}))
	require.NoError(t, err)

	require.Equal(t, "Initiate go.mod\nCreate 'somepkg/typicalw'\n", output.String())
}

func TestSetup_WithGomodFlag_Error(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: "go mod init somepkg",
			ErrorBytes:  []byte("error-message"),
			ReturnError: errors.New("some-error"),
		},
	})
	defer unpatch(t)

	os.Mkdir("somepkg", 0777)
	defer os.RemoveAll("somepkg")

	err := app.Setup(cliContext([]string{
		"-project-pkg=somepkg",
		"-go-mod",
	}))
	require.EqualError(t, err, "some-error: error-message")

	require.Equal(t, "Initiate go.mod\n", output.String())
}

func TestSetup_WithGomodFlag_MissingProjectPkg(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	os.Mkdir("somepkg", 0777)
	defer os.RemoveAll("somepkg")

	err := app.Setup(cliContext([]string{"-go-mod"}))
	require.EqualError(t, err, "project-pkg is empty")

	require.Equal(t, "Initiate go.mod\n", output.String())
}

func TestSetup_WithNewFlag(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	err := app.Setup(cliContext([]string{
		"-project-pkg=somepkg1",
		"-new",
	}))
	require.NoError(t, err)
	defer os.RemoveAll("somepkg1")

	require.Equal(t, `Create 'somepkg1/cmd/somepkg1/main.go'
Create 'somepkg1/internal/app/start.go'
Create 'somepkg1/internal/generated/typical/doc.go'
Create 'somepkg1/tools/typical-build/typical-build.go'
Create 'somepkg1/typicalw'
`, output.String())

	b, _ := ioutil.ReadFile("somepkg1/cmd/somepkg1/main.go")
	require.Equal(t, `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"somepkg1/internal/app"
	_ "somepkg1/internal/generated/typical"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(app.Start)
}
`, string(b))

	b, _ = ioutil.ReadFile("somepkg1/internal/app/start.go")
	require.Equal(t, `package app

import (
	"fmt"
)

// Start app
func Start() {
	fmt.Println("Start app")
	// TODO: 
}
`, string(b))

	b, _ = ioutil.ReadFile("somepkg1/internal/generated/typical/doc.go")
	require.Equal(t, `// Package generated contain generated code from annotate
package generated
`, string(b))

	b, _ = ioutil.ReadFile("somepkg1/tools/typical-build/typical-build.go")
	require.Equal(t, `package main

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	AppName:    "somepkg1",
	AppVersion: "0.0.1",
	AppLayouts: []string{"internal", "pkg"},

	Cmds: []typgo.Cmd{
		// annotate
		&typast.AnnotateCmd{
			Annotators: []typast.Annotator{
				&typapp.CtorAnnotation{},
				&typapp.DtorAnnotation{},
			},
		},
		// compile
		&typgo.CompileCmd{
			Action: &typgo.StdCompile{},
		},
		// run
		&typgo.RunCmd{
			Before: typgo.BuildCmdRuns{"annotate", "compile"},
			Action: &typgo.StdRun{},
		},
		// clean
		&typgo.CleanCmd{
			Action: &typgo.StdClean{},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
`, string(b))
}
