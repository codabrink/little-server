package main

import (
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"./src/server/util"
)

var fileCache = make(map[string][]byte)
var serveDirs = []string{"dst", "src/img"}

func main() {
	makeDirectories()
	util.SpawnWebpack()
	startHttp()
}

func makeDirectories() {
	os.MkdirAll("dst/assets", os.ModePerm)
}

func startHttp() {
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
	if (len(content) == 0) {
		content = scanServe(path)
		if os.Getenv("GO_ENV") == "production" {fileCache[path] = content}
	}
	w.Write(content)
}

func scanServe(path string) []byte {
	var content []byte
	for _, dir := range serveDirs {
		var err error
		content, err = ioutil.ReadFile(filepath.Join(dir, path))
		if err == nil {break}
	}
	if (len(content) == 0) {log.Println(util.Concat("File ", path, " not found."))}
	return content
}
