package main

import (
	"flag"
	"fmt"
	"path/filepath"

	git "github.com/libgit2/git2go/v31"
)

func main() {
	oidPtr := flag.String("oid", "", "oid to search")
	refPtr := flag.String("ref", "", "ref to search")

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: hello-go --oid --ref [dir]")
		return
	}

	dir, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Println("dir is invalid")
		return
	}
	oid := *oidPtr
	ref := *refPtr

	if oid != "" && ref == "" {
		cmdOid(oid, dir)
	} else if ref != "" && oid == "" {
		cmdRef(ref, dir)
	} else if ref != "" && oid != "" {
		cmdRefOid(ref, oid, dir)
	} else {
		fmt.Println("Need either oid or ref")
		return
	}
}

func cmdOid(oid, dir string) {
	fmt.Printf("Looking for oid %s in %s\n", oid, dir)
	found, err := hasObject(oid, dir)
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

func cmdRef(ref, dir string) {
	fmt.Printf("Looking for ref %s in %s\n", ref, dir)
	oid, err := ref2Oid(ref, dir)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}
	fmt.Println(oid.String())
}

func cmdRefOid(ref, oid, dir string) {
	fmt.Printf("Comparing ref %s and oid %s in %s\n", ref, oid, dir)
	oid1, err := ref2Oid(ref, dir)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}
	oid2, err := git.NewOid(oid)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}
	if *oid1 == *oid2 {
		fmt.Println("Matched")
	} else {
		fmt.Printf("Not matched. %s %s\n", oid1.String(), oid2.String())
	}
}

func hasObject(oidString, dir string) (bool, error) {
	repo, err := git.OpenRepository(dir)
	if err != nil {
		return false, fmt.Errorf("failed to open dir. Error: %v", err)
	}

	odb, err := repo.Odb()
	if err != nil {
		return false, fmt.Errorf("failed to open odb. Error: %v", err)
	}

	oid, err := git.NewOid(oidString)
	if err != nil {
		return false, fmt.Errorf("failed to decode oid. Error: %v", err)
	}

	return odb.Exists(oid), nil
}

func ref2Oid(refName, dir string) (*git.Oid, error) {
	repo, err := git.OpenRepository(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to open dir. Error: %v", err)
	}

	if !git.ReferenceIsValidName(refName) {
		return nil, fmt.Errorf("ref %s is not valid", refName)
	}

	ref, err := repo.References.Lookup(refName)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup ref. Error: %v", err)
	}

	resolvedRef, err := ref.Resolve()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve ref. Error: %v", err)
	}

	return resolvedRef.Target(), nil
}
