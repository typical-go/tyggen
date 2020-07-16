package typgo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDtorAnnotation_Execute(t *testing.T) {
	target := "some-target"
	defer os.Remove(target)
	dtorAnnot := &typgo.DtorAnnotation{Target: target}
	ctx := &typgo.Context{
		BuildCli: &typgo.BuildCli{
			ASTStore: &typast.ASTStore{
				Annots: []*typast.Annotation{
					{TagName: "dtor", Decl: &typast.Decl{Name: "Clean", Package: "pkg", Type: typast.FuncType}},
				},
			},
		},
	}

	require.NoError(t, dtorAnnot.Execute(ctx))

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, []byte(`package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
)

func init() { 
	typapp.AppendDestructor(
		&typapp.Destructor{Fn: pkg.Clean},
	)
}`), b)

}

func TestDtorAnnotation_GetTarget(t *testing.T) {
	testcases := []struct {
		TestName string
		*typgo.DtorAnnotation
		Context  *typgo.Context
		Expected string
	}{
		{
			TestName:       "initial target is not set",
			DtorAnnotation: &typgo.DtorAnnotation{},
			Context: &typgo.Context{
				BuildCli: &typgo.BuildCli{
					Descriptor: &typgo.Descriptor{Name: "name0"},
				},
			},
			Expected: "cmd/name0/dtor_annotated.go",
		},
		{
			TestName: "initial target is set",
			DtorAnnotation: &typgo.DtorAnnotation{
				Target: "some-target",
			},
			Expected: "some-target",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.GetTarget(tt.Context))
		})
	}
}
