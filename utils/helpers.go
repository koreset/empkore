package utils

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func SetupTemplates(rootDir string) (*template.Template, error) {
	templates := template.New("")
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			templates, err = templates.New(path).Parse(string(file))
		}
		return nil
	})

	if err != nil{
		return nil, err
	}

	return templates, nil
}
