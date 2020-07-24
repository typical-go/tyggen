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

func TestCtorAnnotation_Annotate(t *testing.T) {
	target := "some-target"
	defer os.Remove(target)

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	ctorAnnot := &typapp.CtorAnnotation{Target: target}
	ctx := &typannot.Context{
		Context: &typgo.Context{},
		ASTStore: &typannot.ASTStore{
			Annots: []*typannot.Annot{
				{
					TagName: "@ctor",
					Decl:    &typannot.Decl{Name: "NewObject", Package: "pkg", Type: &typannot.FuncType{}},
				},
				{
					TagName:  "@ctor",
					TagAttrs: `name:"obj2"`,
					Decl:     &typannot.Decl{Name: "NewObject2", Package: "pkg2", Type: &typannot.FuncType{}},
				},
			},
		},
	}

	require.NoError(t, ctorAnnot.Annotate(ctx))

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() { 
	typapp.AppendCtor(
		&typapp.Constructor{Name: "", Fn: pkg.NewObject},
		&typapp.Constructor{Name: "obj2", Fn: pkg2.NewObject2},
	)
}`, string(b))

}

func TestCtorAnnotation_GetTarget(t *testing.T) {
	testcases := []struct {
		TestName string
		*typapp.CtorAnnotation
		Context  *typannot.Context
		Expected string
	}{
		{
			TestName:       "initial target is not set",
			CtorAnnotation: &typapp.CtorAnnotation{},
			Context: &typannot.Context{
				Context: &typgo.Context{
					BuildSys: &typgo.BuildSys{
						Descriptor: &typgo.Descriptor{Name: "name0"},
					},
				},
			},
			Expected: "cmd/name0/ctor_annotated.go",
		},
		{
			TestName: "initial target is set",
			CtorAnnotation: &typapp.CtorAnnotation{
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
