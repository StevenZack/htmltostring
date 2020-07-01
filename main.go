package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/StevenZack/htmltostring/logx"
)

var (
	templateUse = flag.Bool("t", false, "Use template engine")
	dir         = flag.String("dir", "", "Directory")
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	flag.Parse()
	log.Println("started..")

	d, e := os.Getwd()
	if e != nil {
		log.Println("getwd() failed:", e)
		return
	}
	if len(*dir) > 0 {
		d, e = filepath.Abs(*dir)
		if e != nil {
			log.Println(e)
			return
		}
	}
	fmt.Println(d)
	os.RemoveAll(getrpath(d) + "views")
	os.MkdirAll(getrpath(d)+"views", 0755)

	e = rangeRes(getrpath(d))
	if e != nil {
		logx.Error(e)
		return
	}

	getFilelist(getrpath(d))
}
