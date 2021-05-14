package walnut

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
)

func Generate(code string, config *Config) (string, error) {
	ret := ""
	for _, h := range config.Headers {
		ret += h
	}

	ret += code

	for _, f := range config.Footers {
		ret += f
	}

	templ, err := template.New("template").Parse(ret)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewBufferString("")
	err = templ.Execute(buffer, config.Variables)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func GenerateFile(path string, config *Config) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return Generate(string(contents), config)
}
