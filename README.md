# Typical Go

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)
[![codecov](https://codecov.io/gh/typical-go/typical-go/branch/master/graph/badge.svg)](https://codecov.io/gh/typical-go/typical-go)

A Build Tool (+ Framework) for Golang. <https://typical-go.github.io/>

## Introduction

- Framework to Build-Tool  
  Typical-Go provides levels of abstraction to develop your own build-tool. 
- Build-Tool as a framework (BAAF)  
  It is a concept where both build-tool and application utilize the same definition. We no longer see build-tool as a separate beast with the application but rather part of the same living organism. 


## Wrapper

Wrapper is a simple bash script (`typicalw`) to download, compile and run both build-tool and application. 

```bash
./typicalw
```

```
NAME:
   configuration-with-invocation - Build-Tool

USAGE:
   build-tool [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   test, t     Test the project
   run, r      Run the project in local environment
   publish, p  Publish the project
   clean, c    Clean the project
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Descriptor

The unique about Typical-Go is it use go-based descriptor file rather than DSL which is making it easier to understand and maintain. 

It should be defined at `typical/descriptor.go` with variable name `Descriptor`
```go 
var Descriptor = typgo.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	EntryPoint: server.Main,

	Configurer: server.Configuration(),

	Build: &typgo.StdBuild{},

	Layouts: []string{
		"server",
	},
}
```



## Typical Tmp

Typical-tmp is an important folder that contains the build-tool mechanisms. By default, it is located in `.typical-tmp` and can be changed by hacking/editing the `typicalw` script.

Since the typical-go project is still undergoing development, maybe there is some breaking change and deleting typical-tmp can solve the issue since it will be healed by itself.


## Examples

This repo contain both library, examples and wrapper source-code. The wrapper itself using Typical-Go as its build-tool which is an excellent example.
- [x] [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [Configuration With Invocation](https://github.com/typical-go/typical-go/tree/master/examples/configuration-with-invocation)
- [x] [Simple Additional Task](https://github.com/typical-go/typical-go/tree/master/examples/simple-additional-task)
- [x] [Provide Constructor](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor)
- [x] [Generate Mock](https://github.com/typical-go/typical-go/tree/master/examples/generate-mock)
- [x] [Generate Docker-Compose](https://github.com/typical-go/typical-go/tree/master/examples/generate-docker-compose)
- [x] [Serve React Demo](https://github.com/typical-go/typical-go/tree/master/examples/serve-react-demo)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
