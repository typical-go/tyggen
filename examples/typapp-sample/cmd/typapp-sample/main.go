package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"

	"github.com/typical-go/typical-go/examples/typapp-sample/internal/app"
	_ "github.com/typical-go/typical-go/examples/typapp-sample/internal/generated/ctor"
)

func main() {
	// NOTE: ProjectName and ProjectVersion passed from descriptor in "tools/typical-build" when gobuild
	fmt.Printf("%s %s\n", typgo.ProjectName, typgo.ProjectVersion)

	startFn := app.Start     // What to do when start
	shutdownFn := app.Stop   // What to do when shutdown
	exitSigs := []os.Signal{ // Exit Signals that trigger to close application
		syscall.SIGTERM,
		syscall.SIGINT,
	}

	if err := typapp.StartApp(startFn, shutdownFn, exitSigs...); err != nil {
		log.Fatal(err)
	}
}
