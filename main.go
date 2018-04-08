package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var outfile string

func init() {
	const (
		defaultOutfile = "output"
		usage          = "Output file"
	)

	flag.StringVar(&outfile, "o", defaultOutfile, usage+" (shorthand)")
	flag.StringVar(&outfile, "output", defaultOutfile, usage)
}

func main() {
	flag.Parse()

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("This command is designed to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat file.txt | wrtf -o <outfile>")
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err == nil {
		ioutil.WriteFile(outfile, data, 0644)
	}
}
