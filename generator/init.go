package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type app struct {
	BlntVersion string
	Name        string
	Desc        string
	Module      string
	Database    string
	Port        int
	SecretKey   string
}

var templateURL = "https://raw.githubusercontent.com/bolinette/bolinette-cli/master"

func GenerateHeadlessBolinetteApi(name string, database string) {
	app := app{
		BlntVersion: "0.0.1",
		Name:        name,
		Desc:        name,
		Module:      strings.ToLower(name),
		Database:    database,
		Port:        defaultPortFor(database),
		SecretKey:   generateSecretKey(),
	}
	app.makeFoldersAndEmptyInitPy()
	app.makeFiles()
}

func (app *app) makeFoldersAndEmptyInitPy() {
	var folders = []string{"controllers", "models", "services"}

	ioError(os.Mkdir("env", 0755))

	for _, folder := range folders {
		ioError(os.MkdirAll(fmt.Sprintf("%s/%s", app.Module, folder), 0755))
		createEmptyFile(fmt.Sprintf("%s/%s/__init__.py", app.Module, folder))
	}

	createEmptyFile(fmt.Sprintf("%s/app.py", app.Module))
	createEmptyFile(fmt.Sprintf("%s/seeders.py", app.Module))
}

func (app *app) makeFiles() {
	var apiTemplates = map[string]string{
		fmt.Sprintf("%s/templates/api/.gitignore", templateURL):                          ".gitignore",
		fmt.Sprintf("%s/templates/api/manifest.blnt.yaml", templateURL):                  "manifest.blnt.yaml",
		fmt.Sprintf("%s/templates/api/requirements.txt", templateURL):                    "requirements.txt",
		fmt.Sprintf("%s/templates/api/server/__init__.py", templateURL):                  fmt.Sprintf("%s/__init__.py", app.Module),
		fmt.Sprintf("%s/templates/api/instance/.profile", templateURL):                   "env/.profile",
		fmt.Sprintf("%s/templates/api/instance/env.development.yaml", templateURL):       "env/env.development.yaml",
		fmt.Sprintf("%s/templates/api/instance/env.local.development.yaml", templateURL): "env/env.local.development.yaml",
		fmt.Sprintf("%s/templates/api/instance/env.production.yaml", templateURL):        "env/env.production.yaml",
		fmt.Sprintf("%s/templates/api/instance/init.yaml", templateURL):                  "env/init.yaml",
	}

	for src, dest := range apiTemplates {
		t, err := template.New(src).Parse(downloadFile(src))
		parseTemplateError(err)

		var output bytes.Buffer
		err = t.Execute(&output, app)
		ioError(err)

		err = ioutil.WriteFile(dest, output.Bytes(), 0644)
		ioError(err)
	}
}
