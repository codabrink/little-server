package util

import (
	"bytes"
	"log"
	"io"
	"os"
	"os/exec"
)

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
