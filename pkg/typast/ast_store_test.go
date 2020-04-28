package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestCreateASTStore(t *testing.T) {
	store := typast.CreateASTStore("sample_test.go")
	expectedDecls := []*typast.Decl{
		&typast.Decl{Path: "sample_test.go", Type: typast.Interface, SourceName: "sampleInterface"},
		&typast.Decl{Path: "sample_test.go", Type: typast.Struct, SourceName: "sampleStruct"},
		&typast.Decl{Path: "sample_test.go", Type: typast.Function, SourceName: "sampleFunction"},
	}

	require.Equal(t, len(expectedDecls), len(store.Decls))
	for i, decl := range store.Decls {
		require.True(t, decl.Equal(expectedDecls[i]), decl)
	}

}
