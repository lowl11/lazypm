package pkg_event

import (
	"encoding/json"
	"github.com/lowl11/lazyfile/fileapi"
	"lazypm/src/data/models"
)

const (
	pkgName = "package.json"
)

func (event *Event) read() error {
	pkgFilePath := event.basePath + "/" + pkgName

	if !fileapi.Exist(pkgFilePath) {
		return nil
	}

	fileContent, err := fileapi.Read(pkgFilePath)
	if err != nil {
		return err
	}

	pkg := models.Package{}
	if err = json.Unmarshal(fileContent, &pkg); err != nil {
		return err
	}

	event.pkg = &pkg

	return nil
}
