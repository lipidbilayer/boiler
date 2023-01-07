package statik

import (
	"net/http"

	_ "github.com/lipidbilayer/boiler/lib/files"
	"github.com/rakyll/statik/fs"
)

var statikFS http.FileSystem

type StatikFS struct {
	file *http.FileSystem
}

func New() *StatikFS {
	statikFS, _ = fs.New()
	return &StatikFS{&statikFS}
}

func (s *StatikFS) GetEmbeddedFile() http.FileSystem {
	return *s.file
}
