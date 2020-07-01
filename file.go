package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/StevenZack/tools/strToolkit"

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
		switch filepath.Ext(path) {
		case ".html", ".js", ".css", "tis", ".svg", ".jpeg", ".png", ".ico", "woff", "woff2", "ttf", "ttc", ".glade", ".xml", ".yaml", ".yml":
		default:
			return nil
		}
		name := strToolkit.ToCamelCase(strings.Replace(strings.Replace(f.Name(), ".", "_", -1), "-", "_", -1))
		fo, e := os.OpenFile(root+"views/"+name+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if e != nil {
			log.Println("fo() failed:", e)
			return nil
		}
		defer fo.Close()
		_, e = fo.WriteString(`package views

var Bytes_` + name + " =[]byte{")
		if e != nil {
			log.Println("writeString() failed:", e)
			return nil
		}

		if *templateUse {
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
		} else {
			fi, e := os.OpenFile(path, os.O_RDONLY, 0644)
			if e != nil {
				logx.Error(e)
				return e
			}
			defer fi.Close()
			bs, e := ioutil.ReadAll(fi)
			if e != nil {
				logx.Error(e)
				return e
			}
			fo.WriteString(stringifyBytes(bs))
		}

		fo.WriteString("}\n")
		fmt.Println(root + "views/" + name + ".go")
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func parseFile(fname string) {
	f, e := os.Stat(fname)
	if e != nil {
		log.Println(`os.State error :`, e)
		return
	}
	root, e := os.Getwd()
	if e != nil {
		log.Println(`getwd error :`, e)
		return
	}
	fo, e := os.OpenFile(root+"/views/"+getFirstName(f.Name())+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		log.Println(`openfile error :`, e)
		return
	}
	defer fo.Close()
	_, e = fo.WriteString(`package views

var Str_` + getFirstName(f.Name()) + " =`")
	if e != nil {
		log.Println("writeString() failed:", e)
		return
	}
	fi, e := os.OpenFile(fname, os.O_RDONLY, 0644)
	if e != nil {
		log.Println("fi() failed:", e)
		return
	}
	defer fi.Close()
	_, e = io.Copy(fo, fi)
	if e != nil {
		log.Println("copy() failed:", e)
		return
	}
	fo.WriteString("`\n")
	log.Println(root + "views/" + getFirstName(f.Name()) + ".go")
}
