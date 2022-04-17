package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "/mnt/Drama/更衣人偶坠入爱河.2022",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args[1] = tt.name
			os.Args[2] = "\"[UHA-WINGS][Sono Bisque Doll wa Koi wo Suru][1080p][CHS]\""
			main()
			// todo: check the result
		})
	}
}

func Test_renameFile(t *testing.T) {
	tests := []struct {
		path string
		name string
	}{
		{
			path: "/mnt/Drama/剃须。然后捡到女高中生/",
			name: "剃须。然后捡到女高中生",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, _ := ioutil.ReadDir(tt.path)
			videoFiles := make([]string, 0)
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				filename := file.Name()
				if strings.HasSuffix(filename, ".mp4") {
					videoFiles = append(videoFiles, filename)
				}
			}
			renameFile(videoFiles, tt.name, tt.path)
		})
	}

}
