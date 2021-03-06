package typgen_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	someStructDecl = &typgen.Decl{
		File: typgen.File{
			Path:    "sample_test.go",
			Package: "typgen_test",
		},
		Type: &typgen.StructDecl{
			TypeDecl: typgen.TypeDecl{
				GenDecl: typgen.GenDecl{
					Docs: []string{
						"// sampleStruct",
						"// @tag1",
						"// @tag2 (key1:\"\", key2: \"\", key3:\"value3\")",
					},
				},
				Name: "sampleStruct",
			},
			Fields: []*typgen.Field{
				{
					Names:     []string{"sampleInt"},
					Type:      "int",
					StructTag: reflect.StructTag(`default:"value1"`),
				},
				{
					Names:     []string{"sampleString"},
					Type:      "string",
					StructTag: reflect.StructTag(`default:"value2"`),
				},
			},
		},
	}

	someFunctionDecl2 = &typgen.Decl{
		File: typgen.File{
			Path:    "sample_test.go",
			Package: "typgen_test",
		},
		Type: &typgen.FuncDecl{
			Name:   "sampleFunction2",
			Params: &typgen.FieldList{},
			Docs: []string{
				"// GetWriter to get writer to greet the world",
				"// @ctor",
			},
		},
	}

	someInterface2Decl = &typgen.Decl{
		File: typgen.File{
			Path:    "sample_test.go",
			Package: "typgen_test",
		},
		Type: &typgen.InterfaceDecl{
			TypeDecl: typgen.TypeDecl{
				Name: "sampleInterface2",
				Docs: []string{"// @tag3"},
			},
		},
	}

	someStruct2Decl = &typgen.Decl{
		File: typgen.File{
			Path:    "sample_test.go",
			Package: "typgen_test",
		},
		Type: &typgen.StructDecl{
			TypeDecl: typgen.TypeDecl{
				Name: "sampleStruct2",
				Docs: []string{
					"// sampleStruct2 asdf",
					"// @tag4",
				},
			},
		},
	}
)

func TestGenerator(t *testing.T) {
	Generator := &typgen.Generator{}
	require.Equal(t, &typgo.Task{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate code based on annotation directive ('@')",
		Action:  Generator,
	}, Generator.Task())
}

func TestGenerator_Execute(t *testing.T) {
	var directives typgen.Directives
	action := &typgen.Generator{
		Walker: typgen.FilePaths{"sample_test.go"},
		Processor: &typgen.Annotation{
			ProcessFn: func(c *typgo.Context, d typgen.Directives) error {
				directives = d
				return nil
			},
		},
	}
	require.NoError(t, action.Execute(&typgo.Context{}))
	require.EqualValues(t, typgen.Directives{
		{Decl: someStructDecl, TagName: "@tag1"},
		{Decl: someStructDecl, TagName: "@tag2", TagParam: `key1:"", key2: "", key3:"value3"`},
		{Decl: someFunctionDecl2, TagName: "@ctor"},
		{Decl: someInterface2Decl, TagName: "@tag3"},
		{Decl: someStruct2Decl, TagName: "@tag4"},
	}, directives)
}

func TestGenerator_Error(t *testing.T) {
	testcases := []struct {
		TestName    string
		Generator   *typgen.Generator
		ExpectedErr string
	}{
		{
			Generator: &typgen.Generator{
				Walker: typgen.FilePaths{"bad_file.go"},
			},
			ExpectedErr: "open bad_file.go: no such file or directory",
		},
		{
			Generator: &typgen.Generator{
				Walker: typgen.FilePaths{"sample_test.go"},
				Processor: &typgen.Annotation{
					ProcessFn: func(c *typgo.Context, d typgen.Directives) error {
						return errors.New("some-error")
					},
				},
			},
			ExpectedErr: "some-error",
		},
		{
			Generator:   &typgen.Generator{},
			ExpectedErr: "walker couldn't find any filepath",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			err := tt.Generator.Execute(&typgo.Context{})
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParseRawAnnot(t *testing.T) {
	testcases := []struct {
		TestName         string
		Raw              string
		ExpectedTagName  string
		ExpectedTagAttrs string
	}{
		{
			TestName:        "tag only",
			Raw:             `@tag1`,
			ExpectedTagName: "@tag1",
		},
		{
			TestName:        "tag only with space",
			Raw:             `@tag2 extra1`,
			ExpectedTagName: "@tag2",
		},
		{
			TestName:         "with attribute",
			Raw:              `@tag3("name":"wire1")`,
			ExpectedTagName:  "@tag3",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
		{
			TestName:         "there is space between tagname and attribute",
			Raw:              `@tag4 ("name":"wire1")`,
			ExpectedTagName:  "@tag4",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			tagName, tagAttrs := typgen.ParseRawAnnot(tt.Raw)
			require.Equal(t, tt.ExpectedTagName, tagName)
			require.Equal(t, tt.ExpectedTagAttrs, tagAttrs)
		})
	}
}
