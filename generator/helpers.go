package generator

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func createEmptyFile(name string) {
	d := []byte("")
	ioError(ioutil.WriteFile(name, d, 0644))
}

func downloadFile(url string) string {
	resp, err := http.Get(url)
	httpError(err)
	log.Debug(fmt.Sprintf("Downloading file: %s", url))

	defer resp.Body.Close()

	file, err := ioutil.ReadAll(resp.Body)
	ioError(err)
	return string(file)
}
