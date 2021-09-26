package main

import (
	"fmt"
	"time"

	git "github.com/libgit2/git2go/v32"
)

func main() {
	// test with a real refName
}

func testOidWithOdbCaching(oidString, dir string) {
	oid, err := git.NewOid(oidString)
	if err != nil {
		fmt.Printf("failed to decode oid. Error: %v\n", err)
		return
	}

	repo, err := git.OpenRepository(dir)
	if err != nil {
		fmt.Printf("failed to open dir. Error: %v\n", err)
		return
	}

	odb, err := repo.Odb()
	if err != nil {
		fmt.Printf("failed to open odb. Error: %v\n", err)
		return
	}

	for {
		if odb.Exists(oid) {
			fmt.Printf("oid %s exists\n", oidString)
		} else {
			fmt.Printf("oid %s does not exists\n", oidString)
		}
		time.Sleep(2 * time.Second)
	}
}

func testOidWithoutOdbCaching(oidString, dir string) {
	oid, err := git.NewOid(oidString)
	if err != nil {
		fmt.Printf("failed to decode oid. Error: %v\n", err)
		return
	}

	repo, err := git.OpenRepository(dir)
	if err != nil {
		fmt.Printf("failed to open dir. Error: %v\n", err)
		return
	}

	for {
		odb, err := repo.Odb()
		if err != nil {
			fmt.Printf("failed to open odb. Error: %v\n", err)
			return
		}
		if odb.Exists(oid) {
			fmt.Printf("oid %s exists\n", oidString)
		} else {
			fmt.Printf("oid %s does not exists\n", oidString)
		}
		time.Sleep(2 * time.Second)
	}
}

func testRef(refName, dir string) {
	repo, err := git.OpenRepository(dir)
	if err != nil {
		fmt.Printf("failed to open dir. Error: %v\n", err)
		return
	}

	for {
		if !git.ReferenceIsValidName(refName) {
			fmt.Printf("ref %s is not valid\n", refName)
		} else {
			ref, err := repo.References.Lookup(refName)
			if err != nil {
				fmt.Printf("failed to lookup ref. Error: %v\n", err)
			} else {
				resolvedRef, err := ref.Resolve()
				if err != nil {
					fmt.Printf("failed to resolve ref. Error: %v\n", err)
				} else {
					fmt.Printf("Resolved ref %s to %s\n", refName, resolvedRef.Target().String())
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}
