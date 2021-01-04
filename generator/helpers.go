package generator

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func makeFolders(folders []string, path string, files []string) {
	for _, folder := range folders {
		ioError(os.MkdirAll(fmt.Sprintf("%s/%s", path, folder), 0755))
		if files != nil {
			for _, file := range files {
				createEmptyFile(fmt.Sprintf("%s/%s/%s", path, folder, file))
			}
		}
	}
}

func generatePassword() string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().Unix()))
	return fmt.Sprintf("%x", md5.Sum(b))
}

func generateSecretKey() string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().Unix()))
	return fmt.Sprintf("%x", sha256.Sum224(b))
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

func allToLower(strs []string) []string {
	for i, str := range strs {
		strs[i] = strings.ToLower(str)
	}
	return strs
}
