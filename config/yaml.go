package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Yamlfile interface {
	Get(obj interface{}) error
	Save(obj interface{}) error
}

type yamlfile struct {
	filePath string
	fileName string
}

func newYamlfile(filePath, fileName string) Yamlfile {
	//fmt.Sprintf(configPathFormat, os.Getenv("HOME")), configFileName)
	return &yamlfile{
		filePath: filePath,
		fileName: fileName,
	}
}

func (r *yamlfile) Get(out interface{}) error {
	filename := r.filePath + r.fileName
	strFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(strFile, out); err != nil {
		return err
	}
	return nil
}

func (r *yamlfile) Save(out interface{}) error {
	byteFile, err := yaml.Marshal(out)
	if err != nil {
		return err
	}

	// create dir if doesn't exist
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		if err := os.Mkdir(r.filePath, os.ModePerm); err != nil {
			return err
		}
	}

	// Save file
	filename := r.filePath + r.fileName
	return ioutil.WriteFile(filename, byteFile, 0644)
}
