package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func getrpath(p string) string {
	if p[len(p)-1:] == "/" {
		return p
	}
	return p + "/"
}
func isHiddenDir(s string) bool {
	if s[:1] == "." {
		return true
	}
	if strings.Contains(s, "/.") {
		return true
	}
	return false
}
func getFirstName(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "." {
			return s[:i]
		}
	}
	s = strings.Replace(s, "-", "_", -1)
	return s
}

func readFile(fname string) (string, error) {
	f, e := os.OpenFile(fname, os.O_RDONLY, 0644)
	if e != nil {
		log.Println(e)
		return "", e
	}
	defer f.Close()

	b, e := ioutil.ReadAll(f)
	if e != nil {
		log.Println(e)
		return "", e
	}
	return string(b), nil
}

func getRelativePath(path string) (string, error) {
	wd, e := os.Getwd()
	if e != nil {
		log.Println(e)
		return "", e
	}

	if !strings.HasPrefix(path, wd) {
		return "", errors.New("path is not in wd:" + path)
	}

	return path[len(getrpath(wd)):], nil
}

func stringifyBytes(bs []byte) string {
	builder := &strings.Builder{}
	for _, b := range bs {
		builder.WriteString(strconv.FormatInt(int64(b), 10) + ",")
	}
	return strings.TrimSuffix(builder.String(), ",")
}
