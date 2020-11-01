package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type app struct {
	Name     string
	Database string
}

var apiTemplateURL = "https://github.com/bolinette/bolinette-cli/tree/master/templates/api"

func GenerateHeadlessBolinetteApi(name string, database string) {
	app := app{
		Name:     name,
		Database: database,
	}
	app.makeFoldersAndEmptyInitPy()
	app.makeFiles()
}

func (app *app) makeFoldersAndEmptyInitPy() {
	var folders = []string{"controllers", "models", "services"}

	ioError(os.Mkdir("env", 0755))

	for _, folder := range folders {
		ioError(os.MkdirAll(fmt.Sprintf("%s/%s", app.Name, folder), 0755))
		createEmptyFile(fmt.Sprintf("%s/%s/__init__.py", app.Name, folder))
	}
}

func (app *app) makeFiles() {
	var apiTemplates = map[string]string{
		fmt.Sprintf("%s/templates/api/.gitignore", apiTemplateURL):                          ".gitignore",
		fmt.Sprintf("%s/templates/api/manifest.blnt.yaml", apiTemplateURL):                  "manifest.blnt.yaml",
		fmt.Sprintf("%s/templates/api/requirements.txt", apiTemplateURL):                    "requirements.txt",
		fmt.Sprintf("%s/templates/api/server/__init__.py", apiTemplateURL):                  fmt.Sprintf("%s/__init__.py", app.Name),
		fmt.Sprintf("%s/templates/api/server/app.py", apiTemplateURL):                       fmt.Sprintf("%s/app.py", app.Name),
		fmt.Sprintf("%s/templates/api/server/seeders.py", apiTemplateURL):                   fmt.Sprintf("%s/seeders.py", app.Name),
		fmt.Sprintf("%s/templates/api/instance/.profile", apiTemplateURL):                   "env/.profile",
		fmt.Sprintf("%s/templates/api/instance/env.development.yaml", apiTemplateURL):       "env/env.development.yaml",
		fmt.Sprintf("%s/templates/api/instance/env.local.development.yaml", apiTemplateURL): "env/env.local.development.yaml",
		fmt.Sprintf("%s/templates/api/instance/env.production.yaml", apiTemplateURL):        "env/env.production.yaml",
		fmt.Sprintf("%s/templates/api/instance/init.yaml", apiTemplateURL):                  "env/init.yaml",
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
