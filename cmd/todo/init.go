package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// what this shall do !!
/*
create the folder named .doit inside the parent dir
now inside
.doit/
it needs to have doit.db (This is our main db where everythign happens)
That's all for now
*/

const (
	RootDir = ".doit"
	DirPerm = 0755

// FilePerm   = 0644
// ConfigFile = "config.json"
)

func initProject() {
	fmt.Println("Init Called,  ")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	doitPath := filepath.Join(pwd, RootDir)
	fmt.Println(doitPath)

	// 1. Check if .doit already exists
	if _, err := os.Stat(doitPath); err == nil {
		fmt.Errorf("already a doit repository (directory %s exists)", RootDir)
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Errorf("failed to check status of %s: %w", RootDir, err)
	}

	// creatign .doit with 0755 permision
	if err := os.MkdirAll(doitPath, DirPerm); err != nil {
		fmt.Errorf("failed to create directory %s: %w", doitPath, err)
	}
}
