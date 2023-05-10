package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type BasicParams struct {
	NumRows  int
	Filename string
}

func BasicCzech() (BasicParams, error) {
	var params BasicParams

	name := filepath.Base(os.Args[0])

	// Define and parse the custom flagset
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Způsob použití %s\n", name)
		fs.PrintDefaults()
	}

	// Define and parse the flags
	fs.IntVar(&params.NumRows, "n", 0, "Počet záznamů k vygenerování")
	fs.StringVar(&params.Filename, "o", "", "Název souboru, který chcete vygenerovat")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// Check that the flags were set
	if params.NumRows == 0 || params.Filename == "" {
		fs.Usage()

		err := fmt.Errorf("oba parametry -n a -o jsou vyžadovány")
		fmt.Fprintf(os.Stderr, "\n\nCHYBA: %s\n", err)
		return params, err
	}

	return params, nil
}
