package main

import (
	"flag"
	"log"
	"os"
)

var inDir string
var outDir string

func init() {
	flag.StringVar(&inDir, "i", "in", "input dir")
	flag.StringVar(&outDir, "o", "out", "output dir")
	// flag.Parse()
}

func main() {
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	parseDir(inDir)
}
