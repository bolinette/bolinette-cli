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

	ioError(os.MkdirAll(fmt.Sprintf("%s/env", app.Name), 0755))

	for _, folder := range folders {
		ioError(os.MkdirAll(fmt.Sprintf("%s/src/%s", app.Name, folder), 0755))
		createEmptyFile(fmt.Sprintf("%s/src/%s/__init__.py", app.Name, folder))
	}

	createEmptyFile(fmt.Sprintf("%s/src/app.py", app.Name))
	createEmptyFile(fmt.Sprintf("%s/src/seeders.py", app.Name))
}

func (app *app) makeFiles() {
	var apiTemplates = map[string]string{
		fmt.Sprintf("%s/templates/api/.gitignore", templateURL):                          fmt.Sprintf("%s/.gitignore", app.Name),
		fmt.Sprintf("%s/templates/api/manifest.blnt.yaml", templateURL):                  fmt.Sprintf("%s/manifest.blnt.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/requirements.txt", templateURL):                    fmt.Sprintf("%s/requirements.txt", app.Name),
		fmt.Sprintf("%s/templates/api/server/__init__.py", templateURL):                  fmt.Sprintf("%s/src/__init__.py", app.Name),
		fmt.Sprintf("%s/templates/api/instance/.profile", templateURL):                   fmt.Sprintf("%s/env/.profile", app.Name),
		fmt.Sprintf("%s/templates/api/instance/env.development.yaml", templateURL):       fmt.Sprintf("%s/env/env.development.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/instance/env.local.development.yaml", templateURL): fmt.Sprintf("%s/env/env.local.development.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/instance/env.production.yaml", templateURL):        fmt.Sprintf("%s/env/env.production.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/instance/init.yaml", templateURL):                  fmt.Sprintf("%s/env/init.yaml", app.Name),
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
