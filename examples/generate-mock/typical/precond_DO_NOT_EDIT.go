package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/examples/generate-mock/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() {
	typapp.Provide(
		typapp.NewConstructor("", helloworld.GetWriter),
		typapp.NewConstructor("", helloworld.NewGreeter),
	)
}
