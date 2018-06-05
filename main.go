package main

import (
	"fmt"
	"github.com/StevenZack/tools"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getFilelist(root string) {
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if isHiddenDir(path) {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		if !tools.EndsWith(f.Name(), ".html") && !tools.EndsWith(f.Name(), ".js") && !tools.EndsWith(f.Name(), ".css") {
			// fmt.Println("skip ", f.Name())
			return nil
		}
		fo, e := os.OpenFile(root+"views/"+getFirstName(f.Name())+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if e != nil {
			fmt.Println("fo() failed:", e)
			return nil
		}
		defer fo.Close()
		_, e = fo.WriteString(`package views

var Str_` + getFirstName(f.Name()) + " =`")
		if e != nil {
			fmt.Println("writeString() failed:", e)
			return nil
		}
		fi, e := os.OpenFile(path, os.O_RDONLY, 0644)
		if e != nil {
			fmt.Println("fi() failed:", e)
			return nil
		}
		defer fi.Close()
		_, e = io.Copy(fo, fi)
		if e != nil {
			fmt.Println("copy() failed:", e)
			return nil
		}
		fo.WriteString("`\n")
		fmt.Println(root + "views/" + getFirstName(f.Name()) + ".go")
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	d, e := os.Getwd()
	if e != nil {
		fmt.Println("getwd() failed:", e)
		return
	}
	os.RemoveAll(getrpath(d) + "views")
	os.MkdirAll(getrpath(d)+"views", 0755)
	getFilelist(getrpath(d))
}
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
	return s
}
