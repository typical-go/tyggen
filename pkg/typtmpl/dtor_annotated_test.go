package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestDtorAnnotated(t *testing.T) {
	typtmpl.TestTemplate(t, []typtmpl.TestCase{
		{
			TestName: "common constructor",
			Template: &typtmpl.DtorAnnotated{
				Package: "main",
				Imports: []string{"pkg1", "pkg2"},
				Dtors: []*typtmpl.Dtor{
					{Def: "pkg1.NewFunction1"},
				},
			},
			Expected: `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"pkg1"
	"pkg2"
)

func init() { 
	typapp.Destroy(
		&typapp.Destructor{Fn: pkg1.NewFunction1},
	)
}`,
		},
	})
}
