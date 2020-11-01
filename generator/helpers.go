package generator

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func generateSecretKey() string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().Unix()))
	return fmt.Sprintf("%x", sha256.Sum224(b))
}

func defaultPortFor(database string) int {
	switch database {
	case "SQLITE":
		return 5000
	case "MySql":
		return 3306
	case "PostgreSQL":
		return 5432
	default:
		return 0
	}
}

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
