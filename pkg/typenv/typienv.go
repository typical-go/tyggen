package typenv

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

var (
	appVar        = EnvVar{"TYPICAL_APP", "app"}
	buildToolVar  = EnvVar{"TYPICAL_BUILD_TOOL", "build-tool"}
	prebuilderVar = EnvVar{"TYPICAL_PREBUILDER", "pre-builder"}
	binVar        = EnvVar{"TYPICAL_BIN", "bin"}
	cmdVar        = EnvVar{"TYPICAL_CMD", "cmd"}
	mockVar       = EnvVar{"TYPICAL_MOCK", "mock"}
	releaseVar    = EnvVar{"TYPICAL_RELEASE", "release"}
	dependencyVar = EnvVar{"TYPICAL_DEPENDENCY", "dependency"}
	metadataVar   = EnvVar{"TYPICAL_METADATA", ".typical-metadata"}
	readmeVar     = EnvVar{"TYPICAL_README", "README.md"}
)

var (
	BuildTool  *applicationFolder
	Prebuilder *applicationFolder
	Dependency *applicationFolder
	Bin        string
	Cmd        string
	Metadata   string
	Mock       string
	Release    string
	AppName    string
	Readme     string
)

type applicationFolder struct {
	Package string
	SrcPath string
	BinPath string
}

func init() {
	AppName = appVar.Value()
	Cmd = cmdVar.Value()
	Bin = binVar.Value()
	Metadata = metadataVar.Value()
	buildTool := buildToolVar.Value()
	prebuilder := prebuilderVar.Value()
	dependency := dependencyVar.Value()
	BuildTool = &applicationFolder{
		Package: "main",
		SrcPath: Cmd + "/" + buildTool,
		BinPath: Bin + "/" + buildTool,
	}
	Dependency = &applicationFolder{
		Package: dependency,
		SrcPath: "internal/" + dependency,
	}
	Prebuilder = &applicationFolder{
		Package: "main",
		SrcPath: Cmd + "/" + prebuilder,
		BinPath: Bin + "/" + prebuilder,
	}
	Mock = mockVar.Value()
	Release = releaseVar.Value()
	Readme = readmeVar.Value()
}

// AppMain return main package of application
func AppMain(name string) string {
	return fmt.Sprintf("%s/%s", Cmd, strcase.ToKebab(name))
}

// AppBin return bin path of application
func AppBin(name string) string {
	return fmt.Sprintf("%s/%s", Bin, strcase.ToKebab(name))
}
