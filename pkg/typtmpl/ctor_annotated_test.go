package typtmpl_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestCtorGenerated(t *testing.T) {
	typtmpl.TestTemplate(t, []typtmpl.TestCase{
		{
			TestName: "common constructor",
			Template: &typtmpl.CtorAnnotated{
				Package: "main",
				Imports: []string{"pkg1", "pkg2"},
				Ctors: []*typtmpl.Ctor{
					{Name: "", Def: "pkg1.NewFunction1"},
					{Name: "", Def: "pkg2.NewFunction2"},
				},
			},
			Expected: `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"pkg1"
	"pkg2"
)

func init() { 
	typapp.AppendCtor(
		&typapp.Constructor{Name: "", Fn: pkg1.NewFunction1},
		&typapp.Constructor{Name: "", Fn: pkg2.NewFunction2},
	)
}`,
		},
	})
}

func TestCreateCtor(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		Expected    *typtmpl.Ctor
		ExpectedErr string
	}{
		{
			Annot: &typannot.Annot{
				Decl: &typannot.Decl{
					Package: "pkg",
					Name:    "name",
				},
			},
			Expected: &typtmpl.Ctor{Name: "", Def: "pkg.name"},
		},
		{
			Annot: &typannot.Annot{
				TagAttrs: []byte(`{"name":"some-name"}`),
				Decl: &typannot.Decl{
					Package: "pkg",
					Name:    "name",
				},
			},
			Expected: &typtmpl.Ctor{Name: "some-name", Def: "pkg.name"},
		},
		{
			Annot: &typannot.Annot{
				TagAttrs: []byte(`{bad-attributes`),
				Decl: &typannot.Decl{
					Package: "pkg",
					Name:    "name",
				},
			},
			ExpectedErr: "name: invalid character 'b' looking for beginning of object key string",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctor, err := typtmpl.CreateCtor(tt.Annot)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.Expected, ctor)
			}
		})
	}
}
