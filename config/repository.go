package config

import(
	"fmt"
	"github.com/mitchellh/go-homedir"
)

const baseDirFormat = "%s/.%s/"

type Config interface {
	Filepath(path string)
	Filename(name string)
}

type repository struct {
	filename string
	filepath string
}

func New(name string) Config {
	homepath, _ := homedir.Dir()
	return &repository{
		filename:  "config.yml",
		filepath: fmt.Sprintf(baseDirFormat,homepath, name),
	}
}

func (r *repository) Filepath(path string) {
	r.filepath = path
}

func (r *repository) Filename(name string) {
	r.filename = name
}



