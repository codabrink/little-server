package util

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var fileCache    = make(map[string][]byte)
var serveDirs    = []string{"dst", "src/img"}
var contentTypes = map[string]string{"svg": "image/svg+xml"}

func StartHttp() {
	http.HandleFunc("/", serveDst)
	http.ListenAndServe(":3000", nil)
}

func serveDst(w http.ResponseWriter, r *http.Request) {
	var path string
	if r.URL.Path == "/" {
		path = "/index.html"
	} else {
		path = r.URL.Path
	}

	content := fileCache[path]
	if len(content) == 0 {
		path, content = scanFolders(path)
		if os.Getenv("GO_ENV") == "production" {fileCache[path] = content}
	}
	setContentType(path, w)
	w.Write(content)
	log.Println("Served", path)
}

func setContentType(path string, w http.ResponseWriter) {
	contentType := contentTypes[string(path[len(path)-3:])]
	if len(contentType) == 0 {return}
	w.Header().Set("Content-Type", contentType)
}

func scanFolders(path string) (string, []byte) {
	var content []byte
	var readDir string
	for _, dir := range serveDirs {
		var err error
		readDir = filepath.Join(dir, path)
		content, err = ioutil.ReadFile(readDir)
		if err == nil {break}
	}

	if (len(content) == 0) {log.Println("File ", path, " not found.")}
	return readDir, content
}
