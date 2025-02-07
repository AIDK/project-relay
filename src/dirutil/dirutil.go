package dirutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// Ensure ensures that all directories in the provided slice exits.
// It returns an error if any directory creation fails.
func Ensure(dirs []string, recreate bool) error {

	if recreate {
		for _, dir := range dirs {
			cleanedDir := filepath.Clean(dir)
			fmt.Printf("removing directory (and contents if any): %s\n", cleanedDir)

			// remove all contents within the directory first if it exists
			err := filepath.Walk(cleanedDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// dont remove the root directory
				if path != cleanedDir {
					err := os.RemoveAll(path)
					if err != nil {
						fmt.Printf("error removing %s: %v\n", path, err)
						return err
					}
				}
				return nil
			})

			if err != nil {
				fmt.Printf("error walking directory: %v\n", err)
			}

			// remove the directory itself
			err = os.RemoveAll(cleanedDir)

			// ignore directory if it doesnt exist
			if err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("failed to remove directory %s: %w", cleanedDir, err)
			}

			fmt.Printf("directory %s removed.\n", cleanedDir)
		}
	}

	for _, dir := range dirs {
		cleanedDir := filepath.Clean(dir)

		if _, err := os.Stat(cleanedDir); os.IsNotExist(err) {
			fmt.Printf("Creating directory: %s\n", cleanedDir)
			err := os.MkdirAll(cleanedDir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", cleanedDir, err)
			}

			fmt.Printf("Directory %s created successfully.\n", cleanedDir)

		} else if err != nil {
			return fmt.Errorf("failed to check directory %s: %w", cleanedDir, err)

		} else {
			fmt.Printf("Directory %s already exists.\n", cleanedDir)
		}
	}

	return nil
}

func FileCount(dir string) (count int, err error) {

	// read directory
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	// get files in directory (exclude sub directories)
	files := 0
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			files++
		}
	}

	return files, nil
}
