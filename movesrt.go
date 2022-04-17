package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
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
	var dramaName string
	if len(os.Args) == 3 {
		dramaName = os.Args[2]
	}
	rootPath := os.Args[1]
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Printf("error reading directory: %v", err)
		os.Exit(1)
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
	if dramaName == "" {
		dramaName = path.Base(rootPath)
	}
	dramaName = renameFile(videoFiles, dramaName, rootPath)
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
			// log.Printf("error reading directory: %v\n", err)
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

// 修改文件名，方便Jellyfin解析
func renameFile(files []string, name string, rootPath string) string {
	name = strings.ReplaceAll(name, "】【", ".")
	name = strings.ReplaceAll(name, "】", ".")
	name = strings.ReplaceAll(name, "【", "")
	name = strings.ReplaceAll(name, "][", ".")
	name = strings.ReplaceAll(name, "]", ".")
	name = strings.ReplaceAll(name, "[", "")
	name = strings.ReplaceAll(name, "\"", "")
	log.Printf("drama name: %s\n", name)
	var indexName string
	for index, file := range files {
		if strings.HasSuffix(name, ".") {
			indexName = "E" + strconv.Itoa(index+1)
		} else {
			indexName = ".E" + strconv.Itoa(index+1)
		}
		filename := rootPath + "/" + name + indexName + ".mp4"
		oldfilename := rootPath + "/" + file + ".mp4"
		err := os.Rename(oldfilename, filename)
		if nil == err {
			log.Printf("rename %s to %s %s\n", oldfilename, filename, err)
		}
		log.Printf("rename %s to %s\n", oldfilename, filename)
	}
	return name
}
