package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "" {
		os.Exit(1)
	}
	rootPath := os.Args[1]
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		panic(err)
	}
	videoFiles := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		if strings.HasSuffix(filename, ".mp4") {
			filename = strings.TrimSuffix(filename, ".mp4")
			videoFiles = append(videoFiles, filename)
		}
	}
	var subDir string
	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}
		if dir.Name() == "Subs" {
			subDir = dir.Name()
			break
		}
	}
	for _, file := range videoFiles {
		subFiles, err := ioutil.ReadDir(rootPath + "/" + subDir + "/" + file)
		if err != nil {
			panic(err)
		}
		for _, subFile := range subFiles {
			subFilename := subFile.Name()
			if strings.HasPrefix(subFilename, "34_Chinese") {
				os.Rename(rootPath+"/"+subDir+"/"+file+"/"+subFilename, rootPath+"/"+file+".chs.srt")
				break
			}
		}
	}
}
