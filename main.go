package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/StevenZack/htmltostring/logx"
)

var (
	templateUse = flag.Bool("t", false, "Use template engine")
)

func main() {
	flag.Parse()
	fmt.Println("started..")

	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			parseFile(v)
		}
		return
	}
	d, e := os.Getwd()
	if e != nil {
		fmt.Println("getwd() failed:", e)
		return
	}
	os.RemoveAll(getrpath(d) + "views")
	os.MkdirAll(getrpath(d)+"views", 0755)

	e = rangeRes(getrpath(d))
	if e != nil {
		logx.Error(e)
		return
	}

	getFilelist(getrpath(d))
}
