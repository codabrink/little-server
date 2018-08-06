package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"log"
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	SpawnWebpack()
	StartHttp()
}

func p(a ...interface{}) {fmt.Println(a)}

func SpawnWebpack() {
	var stdBuffer bytes.Buffer
	cmd := exec.Command("bash", "-c", "npx webpack --config webpack.config.js")
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw
	err := cmd.Start()
	if err != nil {panic(err)}
	log.Println(stdBuffer.String())
}

var fileCache = make(map[string][]byte)
func serveDist(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("dist")
	if r.URL.Path == "/" {
		buffer.WriteString("/index.html")
	} else {
		buffer.WriteString(r.URL.Path)
	}

	filepath := buffer.String()

	content := fileCache[filepath]
	if (len(content) == 0) {
		var err error
		content, err = ioutil.ReadFile(filepath)
		if err != nil {log.Fatal(err)}
		if os.Getenv("GO_ENV") == "production" {fileCache[filepath] = content}
	}
	w.Write(content)
}

func StartHttp() {
	os.MkdirAll("./dist", os.ModePerm)
	http.HandleFunc("/", serveDist)
	http.ListenAndServe(":3000", nil)
}
