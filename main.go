package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/haccer/subjack/subjack"
)

//go:embed fingerprints.json
var configData string

func main() {
	o := subjack.Options{}

	flag.StringVar(&o.Domain, "d", "", "Domain.")
	flag.StringVar(&o.Wordlist, "w", "", "Path to wordlist.")
	flag.IntVar(&o.Threads, "t", 10, "Number of concurrent threads (Default: 10).")
	flag.IntVar(&o.Timeout, "timeout", 10, "Seconds to wait before connection timeout (Default: 10).")
	flag.BoolVar(&o.Ssl, "ssl", false, "Force HTTPS connections (May increase accuracy (Default: http://).")
	flag.BoolVar(&o.All, "a", false, "Find those hidden gems by sending requests to every URL. (Default: Requests are only sent to URLs with identified CNAMEs).")
	flag.BoolVar(&o.Verbose, "v", false, "Display more information per each request.")
	flag.StringVar(&o.Output, "o", "", "Output results to file (Subjack will write JSON if file ends with '.json').")
	flag.StringVar(&o.Config, "c", "", "Path to configuration file.")
	flag.BoolVar(&o.Manual, "m", false, "Flag the presence of a dead record, but valid CNAME entry.")

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// If no config file is provided, write the embedded config to a temporary file
	if o.Config == "" {
		tempDir := os.TempDir()
		tempFile := filepath.Join(tempDir, "fingerprints.json")

		err := os.WriteFile(tempFile, []byte(configData), 0644)
		if err != nil {
			fmt.Println("Error writing embedded config to temp file:", err)
			os.Exit(1)
		}

		o.Config = tempFile
	}

	subjack.Process(&o)
}
