package typtmpl_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestExecute(t *testing.T) {
	t.Run("WHEN success", func(t *testing.T) {
		var builder strings.Builder
		require.NoError(t, typtmpl.Execute("", "hello {{.Name}}", &data{Name: "world"}, &builder))
		require.Equal(t, "hello world", builder.String())
	})
	t.Run("WHEN error", func(t *testing.T) {
		require.EqualError(t,
			typtmpl.Execute("", "bad-template {{{}", nil, nil),
			"template: :1: unexpected \"{\" in command",
		)
	})
}

type data struct {
	Name string
}

type dummyTemplate struct {
	text string
}

func (s *dummyTemplate) Execute(w io.Writer) (err error) {
	w.Write([]byte(s.text))
	return
}
