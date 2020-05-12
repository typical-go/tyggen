package typgo

import (
	"github.com/typical-go/typical-go/pkg/git"
)

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*Context) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*Context) error
}

// Releaser responsible to release
type Releaser interface {
	Release(*ReleaseContext) (files []string, err error)
}

// Publisher responsible to publish the release to external source
type Publisher interface {
	Publish(*PublishContext) error
}

// Runner responsible to run the project in local environment
type Runner interface {
	Run(c *Context) error
}

// ReleaseContext is context of release
type ReleaseContext struct {
	*Context
	Alpha   bool
	Tag     string
	GitLogs []*git.Log
}

// PublishContext is context of publish
type PublishContext struct {
	*ReleaseContext
	ReleaseFiles []string
}