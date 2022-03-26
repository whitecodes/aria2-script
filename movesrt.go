package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: movesrt <path>")
		os.Exit(1)
	}
	if os.Args[1] == "" {
		log.Println("must specify a path")
		os.Exit(1)
	}
	rootPath := os.Args[1]
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Printf("error reading directory: %v", err)
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

	log.Printf("found %d video files\n", len(videoFiles))
	var subDir string
	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}
		if dir.Name() == "Subs" {
			subDir = dir.Name()
			log.Println("found Subs directory")
			break
		}
	}
	for _, file := range videoFiles {
		subFiles, err := ioutil.ReadDir(rootPath + "/" + subDir + "/" + file)
		if err != nil {
			log.Printf("error reading directory: %v\n", err)
			log.Printf("skipping %s\n", file)
			continue
		}
		foundSub := false
		for _, subFile := range subFiles {
			subFilename := subFile.Name()
			if strings.HasPrefix(subFilename, "34_Chinese") {
				log.Printf("found Chinese subtitle for %s\n", file)
				os.Rename(rootPath+"/"+subDir+"/"+file+"/"+subFilename, rootPath+"/"+file+".chs.srt")
				log.Println("move subtitle file finish")
				foundSub = true
				break
			}
		}
		if !foundSub {
			log.Printf("no Chinese subtitle found for %s\n", file)
		}
	}
}
