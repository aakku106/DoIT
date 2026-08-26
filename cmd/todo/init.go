package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	RootDir = ".doit"
	DirPerm = 0755
)

func initProject() {
	fmt.Println("Init Called...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	doitPath := filepath.Join(pwd, RootDir)

	// 1. Check if .doit already exists
	if _, err := os.Stat(doitPath); err == nil {
		fmt.Fprintf(os.Stderr, "Error: already a doit repository (%s exists)\n", RootDir)
		os.Exit(1)
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Error checking status of %s: %v\n", RootDir, err)
		os.Exit(1)
	}

	// 2. Create .doit folder
	if err := os.MkdirAll(doitPath, DirPerm); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", doitPath, err)
		os.Exit(1)
	}

	// 3. Create empty doit.db file inside .doit
	dbPath := filepath.Join(doitPath, "doit.db")
	dbFile, err := os.Create(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", dbPath, err)
		os.Exit(1)
	}
	dbFile.Close()

	fmt.Printf("Initialized empty doit repository in %s\n", doitPath)
}
