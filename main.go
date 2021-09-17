package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"path/filepath"

	git "github.com/libgit2/git2go/v31"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: hello-go [dir] [oid]")
		return
	}

	dir, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Println("dir is invalid")
		return
	}
	oid := args[1]

	fmt.Printf("Looking for %s in %s\n", oid, dir)
	found, err := hasObject(dir, oid)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}
	if found {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}

func hasObject(dir, oid string) (bool, error) {
	repo, err := git.OpenRepository(dir)
	if err != nil {
		return false, fmt.Errorf("Failed to open dir. Error: %v", err)
	}

	odb, err := repo.Odb()
	if err != nil {
		return false, fmt.Errorf("Failed to open odb. Error: %v", err)
	}

	oidb, err := hex.DecodeString(oid)
	if err != nil {
		return false, fmt.Errorf("Failed to decode oid. Error: %v", err)
	}

	return odb.Exists((*git.Oid)(oidb)), nil
}
