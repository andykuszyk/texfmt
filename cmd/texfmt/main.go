package main

import (
	"log"
	"flag"
	"github.com/andykuszyk/texfmt/internal/formatter"
	"os"
)

func main() {
	width := flag.Int("w", 120, "The line width to format the file to")
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("Usage: texfmt [-w <line-width>] <file-path>")
	}
	file := flag.Arg(0)
	formatted, err := formatter.Format(file, *width)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.WriteString(formatted)
}

