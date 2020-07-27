package typapp_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDtorAnnotation_Execute(t *testing.T) {
	target := "some-target"
	defer os.Remove(target)

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	dtorAnnot := &typapp.DtorAnnotation{Target: target}
	ctx := &typannot.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{},
			},
		},
		ASTStore: &typannot.ASTStore{
			Annots: []*typannot.Annot{
				{TagName: "@dtor", Decl: &typannot.Decl{Name: "Clean", Package: "pkg", Type: &typannot.FuncType{}}},
			},
		},
	}

	require.NoError(t, dtorAnnot.Annotate(ctx))

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() { 
	typapp.AppendDtor(
		&typapp.Destructor{Fn: pkg.Clean},
	)
}`, string(b))

}

func TestDtorAnnotation_Annotate_RemoveTargetWhenNoAnnotation(t *testing.T) {
	target := "some-target"
	defer os.Remove(target)
	ioutil.WriteFile(target, []byte("some-content"), 0777)
	dtorAnnot := &typapp.DtorAnnotation{Target: target}
	ctx := &typannot.Context{
		Context:  &typgo.Context{},
		ASTStore: &typannot.ASTStore{},
	}
	require.NoError(t, dtorAnnot.Annotate(ctx))
	_, err := os.Stat(target)
	require.True(t, os.IsNotExist(err))
}

func TestDtorAnnotation_GetTarget(t *testing.T) {
	testcases := []struct {
		TestName string
		*typapp.DtorAnnotation
		Context  *typannot.Context
		Expected string
	}{
		{
			TestName:       "initial target is not set",
			DtorAnnotation: &typapp.DtorAnnotation{},
			Context: &typannot.Context{
				Context: &typgo.Context{
					BuildSys: &typgo.BuildSys{
						Descriptor: &typgo.Descriptor{Name: "name0"},
					},
				},
			},
			Expected: "cmd/name0/dtor_annotated.go",
		},
		{
			TestName: "initial target is set",
			DtorAnnotation: &typapp.DtorAnnotation{
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
