package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/zjykzk/geektime"
)

var (
	inputDir  string
	outputDir string
)

func init() {
	flag.StringVar(&inputDir, "input", "", "input dir, generated by the downloader")
	flag.StringVar(&outputDir, "output", "", "output dir")
}

func main() {
	if err := validParams(); err != nil {
		fmt.Println(err.Error())
		return
	}

	c, err := geektime.NewM3U8Converter(inputDir, outputDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.ToMP3()
}

func validParams() error {
	flag.Parse()
	if inputDir == "" {
		return errors.New("empty input dir")
	}

	if outputDir == "" {
		return errors.New("empty ouput dir")
	}
	return nil
}
