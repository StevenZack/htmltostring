package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/StevenZack/htmltostring/logx"
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
		if !strings.HasSuffix(f.Name(), ".html") && !strings.HasSuffix(f.Name(), ".js") && !strings.HasSuffix(f.Name(), ".tis") && !strings.HasSuffix(f.Name(), ".css") && !strings.HasSuffix(f.Name(), ".svg") {
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

		input, e := readFile(path)
		if e != nil {
			logx.Error(e)
			return e
		}

		tpl, e := template.New("name").Parse(input)
		if e != nil {
			logx.Error(e)
			return e
		}
		tpl.Execute(fo, resMap)

		fo.WriteString("`\n")
		fmt.Println(root + "views/" + getFirstName(f.Name()) + ".go")
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func parseFile(fname string) {
	f, e := os.Stat(fname)
	if e != nil {
		fmt.Println(`os.State error :`, e)
		return
	}
	root, e := os.Getwd()
	if e != nil {
		fmt.Println(`getwd error :`, e)
		return
	}
	fo, e := os.OpenFile(root+"/views/"+getFirstName(f.Name())+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Println(`openfile error :`, e)
		return
	}
	defer fo.Close()
	_, e = fo.WriteString(`package views

var Str_` + getFirstName(f.Name()) + " =`")
	if e != nil {
		fmt.Println("writeString() failed:", e)
		return
	}
	fi, e := os.OpenFile(fname, os.O_RDONLY, 0644)
	if e != nil {
		fmt.Println("fi() failed:", e)
		return
	}
	defer fi.Close()
	_, e = io.Copy(fo, fi)
	if e != nil {
		fmt.Println("copy() failed:", e)
		return
	}
	fo.WriteString("`\n")
	fmt.Println(root + "views/" + getFirstName(f.Name()) + ".go")
}
