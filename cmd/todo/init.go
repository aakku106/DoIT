package main

import (
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
}
