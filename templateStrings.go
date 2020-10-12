package main

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const stringsDir = "templates/strings"
const defaultStrings = "default"
const variantFileExt = ".yaml"

var templateStrings map[string]map[string]string

func initTemplateStrings() error {
	templateStrings = map[string]map[string]string{}
	variants := []string{defaultStrings}
	for _, blog := range appConfig.Blogs {
		variants = append(variants, blog.Lang)
	}
	for _, variant := range variants {
		variantStrings := map[string]string{}
		fileContent, err := ioutil.ReadFile(path.Join(stringsDir, variant+variantFileExt))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		err = yaml.Unmarshal(fileContent, variantStrings)
		if err != nil {
			return err
		}
		templateStrings[variant] = variantStrings
	}
	return nil
}

func getTemplateStringVariant(lang, name string) (result string) {
	m, ok := templateStrings[lang]
	if !ok {
		m = templateStrings[defaultStrings]
	}
	result, ok = m[name]
	if !ok {
		result = templateStrings[defaultStrings][name]
	}
	return
}