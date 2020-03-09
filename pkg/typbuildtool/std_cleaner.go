package typbuildtool

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// StdCleaner is standard cleaner
type StdCleaner struct {
}

// NewCleaner return new instance of StdCleaner
func NewCleaner() *StdCleaner {
	return &StdCleaner{}
}

// Clean the project
func (*StdCleaner) Clean(c *Context) (err error) {
	removeAllFile(c.BinFolder)
	removeAllFile(c.TempFolder)
	return
}

func removeFile(name string) {
	log.Infof("Remove: %s", name)
	if err := os.Remove(name); err != nil {
		log.Error(err.Error())
	}
}

func removeAllFile(path string) {
	log.Infof("Remove All: %s", path)
	if err := os.RemoveAll(path); err != nil {
		log.Error(err.Error())
	}
}