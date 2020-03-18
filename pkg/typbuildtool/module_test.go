package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestStdBuilder(t *testing.T) {
	t.Run("SHOULD implement Builder", func(t *testing.T) {
		var _ typbuildtool.Builder = typbuildtool.NewModule()
	})
	t.Run("SHOULD implement Cleaner", func(t *testing.T) {
		var _ typbuildtool.Cleaner = typbuildtool.NewModule()
	})
	t.Run("SHOULD implement Mock", func(t *testing.T) {
		var _ typbuildtool.Mocker = typbuildtool.NewModule()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Tester = typbuildtool.NewModule()
	})
	t.Run("SHOULD implement Release", func(t *testing.T) {
		var _ typbuildtool.Releaser = typbuildtool.NewModule()
	})
}

func TestValidate(t *testing.T) {
	testcases := []struct {
		*typbuildtool.Module
		expected string
	}{
		{
			Module:   typbuildtool.NewModule().WithReleaseTargets(),
			expected: "Missing 'Targets'",
		},
		{
			Module:   typbuildtool.NewModule().WithReleaseTargets("invalid-target"),
			expected: "Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.expected, i)
	}
}

func TestStdMocker(t *testing.T) {
	mocker := typbuildtool.NewModule()
	mocker.PutMockTarget(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target1"})
	mocker.PutMockTarget(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target2"})
	mocker.PutMockTarget(&typbuildtool.MockTarget{MockDir: "pkg2", SrcName: "target3"})
	mocker.PutMockTarget(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target4"})
	mocker.PutMockTarget(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target5"})
	mocker.PutMockTarget(&typbuildtool.MockTarget{MockDir: "pkg2", SrcName: "target6"})

	require.Equal(t, map[string][]*typbuildtool.MockTarget{
		"pkg1": []*typbuildtool.MockTarget{
			{MockDir: "pkg1", SrcName: "target1"},
			{MockDir: "pkg1", SrcName: "target2"},
			{MockDir: "pkg1", SrcName: "target4"},
			{MockDir: "pkg1", SrcName: "target5"},
		},
		"pkg2": []*typbuildtool.MockTarget{
			{MockDir: "pkg2", SrcName: "target3"},
			{MockDir: "pkg2", SrcName: "target6"},
		},
	}, mocker.MockTargetMap())
}