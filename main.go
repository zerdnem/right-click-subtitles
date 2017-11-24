package main

import (
	"os"
	"path/filepath"
	"regexp"

	"./lib"
)

//Subtitles ...
type Subtitles struct {
	filenum  int
	filename string
	link     string
}

func checkExt(ext string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

func main() {

	ext := []string{".avi", ".mp4", ".mkv", ".mpg", ".mpeg", ".mov", ".rm", ".vob", ".wmv", ".flv", ".3gp", ".3g2"}

	language := []string{"eng"}

	for _, v := range ext {

		f := checkExt(v)
		if len(f) > 0 {
			lib.SearchSubtitlesInPath(f[0], language)
		}

	}

}
