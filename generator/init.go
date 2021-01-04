package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

type app struct {
	BlntVersion      string
	Name             string
	Desc             string
	Module           string
	Database         string
	DatabasePassword string
	Port             int
	SecretKey        string
}

var templateURL = "https://raw.githubusercontent.com/bolinette/bolinette-cli/master"

func GenerateHeadlessBolinetteApi(name string, database string) {
	app := app{
		BlntVersion:      "0.0.1",
		Name:             name,
		Desc:             name,
		Module:           strings.ToLower(name),
		Database:         strings.ToLower(database),
		DatabasePassword: generatePassword(),
		Port:             5000,
		SecretKey:        generateSecretKey(),
	}
	app.createFoldersAndEmptyPyFiles()
	app.createAPIFilesFromTemplates()
	app.createDockerFilesFromTemplates()
}

func (app *app) createFoldersAndEmptyPyFiles() {
	var srcFolders = []string{"controllers", "models", "services"}
	var blntFolders = []string{"env", "docker"}

	makeFolders([]string{"src"}, app.Name, nil)
	makeFolders(srcFolders, fmt.Sprintf("%s/src", app.Name), []string{"__init__.py"})
	makeFolders(blntFolders, app.Name, nil)
}

func (app *app) createAPIFilesFromTemplates() {
	var apiTemplates = map[string]string{
		fmt.Sprintf("%s/templates/api/.gitignore", templateURL):                          fmt.Sprintf("%s/.gitignore", app.Name),
		fmt.Sprintf("%s/templates/api/manifest.blnt.yaml", templateURL):                  fmt.Sprintf("%s/manifest.blnt.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/requirements.txt", templateURL):                    fmt.Sprintf("%s/requirements.txt", app.Name),
		fmt.Sprintf("%s/templates/api/server.py", templateURL):                           fmt.Sprintf("%s/server.py", app.Name),
		fmt.Sprintf("%s/templates/api/server/__init__.py", templateURL):                  fmt.Sprintf("%s/src/__init__.py", app.Name),
		fmt.Sprintf("%s/templates/api/server/app.py", templateURL):                       fmt.Sprintf("%s/src/app.py", app.Name),
		fmt.Sprintf("%s/templates/api/server/seeders.py", templateURL):                   fmt.Sprintf("%s/src/seeders.py", app.Name),
		fmt.Sprintf("%s/templates/api/instance/.profile", templateURL):                   fmt.Sprintf("%s/env/.profile", app.Name),
		fmt.Sprintf("%s/templates/api/instance/env.development.yaml", templateURL):       fmt.Sprintf("%s/env/env.development.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/instance/env.local.development.yaml", templateURL): fmt.Sprintf("%s/env/env.local.development.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/instance/env.production.yaml", templateURL):        fmt.Sprintf("%s/env/env.production.yaml", app.Name),
		fmt.Sprintf("%s/templates/api/instance/init.yaml", templateURL):                  fmt.Sprintf("%s/env/init.yaml", app.Name),
	}

	app._createFilesFromTemplate(apiTemplates)
}

func (app *app) createDockerFilesFromTemplates() {
	var dockerTemplates = map[string]string{
		fmt.Sprintf("%s/templates/docker/Dockerfile", templateURL):                      fmt.Sprintf("%s/docker/Dockerfile", app.Name),
		fmt.Sprintf("%s/templates/docker/databases/%s.yaml", templateURL, app.Database): fmt.Sprintf("%s/docker/db.yaml", app.Name),
		fmt.Sprintf("%s/templates/docker/app.yaml", templateURL):                        fmt.Sprintf("%s/docker/%s.yaml", app.Name, app.Module),
	}

	app._createFilesFromTemplate(dockerTemplates)
}

func (app *app) _createFilesFromTemplate(templateToFiles map[string]string) {
	for src, dest := range templateToFiles {
		t, err := template.New(src).Parse(downloadFile(src))
		parseTemplateError(err)

		var output bytes.Buffer
		err = t.Execute(&output, app)
		ioError(err)

		err = ioutil.WriteFile(dest, output.Bytes(), 0644)
		ioError(err)
	}
}
