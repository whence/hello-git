package main

import (
	"fmt"
	"os"
	"time"

	git "github.com/libgit2/git2go/v32"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[prog] [test_case] args...")
		return
	}

	switch testCase := os.Args[1]; testCase {
	case "OidWithOdbCaching":
		testOidWithOdbCaching(os.Args[2], os.Args[3])
	case "OidWithoutOdbCaching":
		testOidWithoutOdbCaching(os.Args[2], os.Args[3])
	case "Ref":
		testRef(os.Args[2], os.Args[3])
	default:
		fmt.Println("Unknown test case")
	}
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
