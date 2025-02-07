package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AIDK/project-relay/src/dirutil"
	"github.com/AIDK/project-relay/src/parser"
)

func main() {

	// recreate resets all application directories (cleans out all files)
	recreate := flag.Bool("recreate", false, "recreate directories (delete existing ones)")
	flag.Parse()

	// parse application config file
	tomlFile, err := parser.ParseConfig()
	if err != nil {
		log.Fatalf("error reading TOML file: %v", err)
	}

	// create directories
	err = dirutil.Ensure(tomlFile.Directories, *recreate)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("directory setup complete")
}
