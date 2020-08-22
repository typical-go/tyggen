package typannot_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

var (
	someInterfaceDecl = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.InterfaceDecl{
			TypeDecl: typannot.TypeDecl{Name: "sampleInterface"},
		},
	}

	someStructDecl = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.StructDecl{
			TypeDecl: typannot.TypeDecl{
				GenDecl: typannot.GenDecl{
					Docs: []string{
						"// sampleStruct",
						"// @tag1",
						"// @tag2 (key1:\"\", key2: \"\", key3:\"value3\")",
					},
				},
				Name: "sampleStruct",
			},
			Fields: []*typannot.Field{
				{Name: "sampleInt", Type: "int", StructTag: reflect.StructTag(`default:"value1"`)},
				{Name: "sampleString", Type: "string", StructTag: reflect.StructTag(`default:"value2"`)},
			},
		},
	}

	someFunctionDecl = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.FuncDecl{Name: "sampleFunction"},
	}

	someFunctionDecl2 = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.FuncDecl{
			Name: "sampleFunction2",
			Docs: []string{
				"// GetWriter to get writer to greet the world",
				"// @ctor",
			},
		},
	}

	someInterface2Decl = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.InterfaceDecl{
			TypeDecl: typannot.TypeDecl{
				Name: "sampleInterface2",
				Docs: []string{"// @tag3"},
			},
		},
	}

	someStruct2Decl = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.StructDecl{
			TypeDecl: typannot.TypeDecl{
				Name: "sampleStruct2",
				Docs: []string{
					"// sampleStruct2 asdf",
					"// @tag4",
				},
			},
		},
	}
	someStruct3Decl = &typannot.Decl{
		File: typannot.File{
			Path:    "sample_test.go",
			Package: "typannot_test",
		},
		DeclType: &typannot.StructDecl{
			TypeDecl: typannot.TypeDecl{Name: "sampleStruct3"},
			Fields: []*typannot.Field{
				{Name: "Name", Type: "string"},
				{Name: "Address", Type: "string"},
			},
		},
	}
)

func TestCompile(t *testing.T) {
	summary, err := typannot.Compile("sample_test.go")
	require.NoError(t, err)

	require.EqualValues(t, []*typannot.Decl{
		someInterfaceDecl,
		someStructDecl,
		someFunctionDecl,
		someFunctionDecl2,
		someInterface2Decl,
		someStruct2Decl,
		someStruct3Decl,
	}, summary.Decls)

	// require.EqualValues(t, []*typannot.Annot{
	// 	{
	// 		Decl:    someStructDecl,
	// 		TagName: "@tag1",
	// 	},
	// 	{
	// 		Decl:     someStructDecl,
	// 		TagName:  "@tag2",
	// 		TagParam: `key1:"", key2: "", key3:"value3"`,
	// 	},
	// 	{
	// 		Decl:    someFunctionDecl2,
	// 		TagName: "@ctor",
	// 	},
	// 	{
	// 		Decl:    someInterface2Decl,
	// 		TagName: "@tag3",
	// 	},
	// 	{
	// 		Decl:    someStruct2Decl,
	// 		TagName: "@tag4",
	// 	},
	// }, summary.Annots)
}

func TestCompile_FileNotFound(t *testing.T) {
	_, err := typannot.Compile("not_found.go")
	require.EqualError(t, err, "open not_found.go: no such file or directory")
}

func TestWalk(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
	}()

	dirs, files := typannot.Walk([]string{"pkg", "wrapper"})
	require.Equal(t, []string{"pkg", "pkg/some_lib", "wrapper", "wrapper/some_pkg"}, dirs)
	require.Equal(t, []string{"pkg/some_lib/lib.go", "wrapper/some_pkg/some_file.go"}, files)
}
