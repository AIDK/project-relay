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

	//STEP
	// recreate resets all application directories (cleans out all files)
	recreate := flag.Bool("recreate", false, "recreate directories (delete existing ones)")
	flag.Parse()

	//STEP
	// parse application config file
	tomlFile, err := parser.ParseConfig()
	if err != nil {
		log.Fatalf("error reading TOML file: %v", err)
	}

	//STEP
	// create directories
	err = dirutil.Ensure(tomlFile.Directories, *recreate)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("directory setup complete")

	//STEP
	// get total number of files in root directory (required to determine whether to start processing)
	totalFiles, err := dirutil.FileCount(tomlFile.Directories[0])
	if err != nil {
		log.Fatalf("error getting file count: %v", err)
	}

	fmt.Printf("number of file in directory: %d", totalFiles)

	//STEP
	// process file
	if totalFiles > 0 {
		// continue
	}
}
