package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AIDK/project-relay/src/dirutil"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Directories []string `toml:"directories"`
}

func main() {

	// recreate resets all application directories (cleans out all files)
	recreate := flag.Bool("recreate", false, "recreate directories (delete existing ones)")
	flag.Parse()

	// read the configuration file (TOML)
	tomlFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatalf("error reading TOML file: %v", err)
	}

	// decode TOML file to config struct
	var config Config
	if _, err := toml.Decode(string(tomlFile), &config); err != nil {
		log.Fatalf("error decoding TOML: %v", err)
	}

	// create directories
	err = dirutil.Ensure(config.Directories, *recreate)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("directory setup complete")
}
