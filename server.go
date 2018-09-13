package main

import (
	"os"
	"./src/server/util"
)

func main() {
	makeDirectories()
	util.SpawnWebpack()
	util.StartHttp()
}

func makeDirectories() {
	os.MkdirAll("dst/assets", os.ModePerm)
}
