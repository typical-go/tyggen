package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/example/helloworld/somepkg"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typdep"
)

func init() {
	typapp.AppendConstructor(
		typdep.NewConstructor(somepkg.NewSomeStruct),
	)
}
