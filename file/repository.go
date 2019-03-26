package file

import (
	"errors"
	"io/ioutil"
)

type Repository interface {
	GetDirs(url, filepath string) error
	Update(url, filepath string) error
	Delete(filepath string) error
}



func getFiles(){
	filePath := "./"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		//&& (filepath.Ext(file.Name()) == ".csv" || filepath.Ext(file.Name()) == ".csv.gz")
		if !file.IsDir() {
			options = append(options, file.Name())
		}
	}

	if len(options) == 0 {
		return "", errors.New("No csv files in the configured directory. Run 'csvpub config show'")
	}

}