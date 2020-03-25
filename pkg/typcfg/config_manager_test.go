package typcfg_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewConfigManager(t *testing.T) {
	t.Run("", func(t *testing.T) {
		cfgMngr := typcfg.Create()
		require.Equal(t, "*typcfg.defaultLoader", reflect.TypeOf(cfgMngr.Loader()).String())
	})
	t.Run("SHOULD implement of ConfigManager", func(t *testing.T) {
		var _ typcore.ConfigManager = typcfg.Create()
	})
	t.Run("SHOULD implement of Preconditioner", func(t *testing.T) {
		var _ typbuildtool.Preconditioner = typcfg.Create()
	})
	t.Run("", func(t *testing.T) {
		configuration := typcfg.
			Create(&dummyConfigurer1{}, &dummyConfigurer2{}).
			WithLoader(&dummyLoader{})

		require.IsType(t, &dummyLoader{}, configuration.Loader())

		var b strings.Builder
		require.NoError(t, configuration.Write(&b))
		require.Equal(t, "prefix1_ID=\nprefix1_VOLUME=\nprefix2_TITLE=default-title\nprefix2_CONTENT=default-content\n", b.String())
	})
}

type dummyLoader struct{}

func (*dummyLoader) LoadConfig(string, interface{}) error { return nil }

type dummyConfigurer1 struct{}

func (*dummyConfigurer1) Configure() *typcfg.Configuration {
	return typcfg.NewConfiguration("prefix1", &struct {
		ID     int64 ``
		Volume int   ``
	}{})
}

type dummyConfigurer2 struct{}

func (*dummyConfigurer2) Configure() *typcfg.Configuration {
	return typcfg.NewConfiguration("prefix2", &struct {
		Title   string `default:"default-title"`
		Content string `default:"default-content"`
	}{})
}
