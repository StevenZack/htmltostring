package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var resMap = make(map[string]string)

func rangeRes(root string) error {
	return filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if isHiddenDir(path) {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		if !strings.HasSuffix(f.Name(), ".html") && !strings.HasSuffix(f.Name(), ".js") && !strings.HasSuffix(f.Name(), ".tis") && !strings.HasSuffix(f.Name(), ".css") && !strings.HasSuffix(f.Name(), ".svg") {
			// fmt.Println("skip ", f.Name())
			return nil
		}

		input, e := readFile(path)
		if e != nil {
			log.Println(e)
			return e
		}

		relative, e := getRelativePath(path)
		if e != nil {
			log.Println(e)
			return e
		}

		resMap[relative] = input
		return nil
	})
}
