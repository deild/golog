// +build mage

// nolint
package main

import (
	"errors"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(GetDeps)
	return sh.RunV("vgo", "build", "-o", "golog", ".")
}

// Manage your deps, or running package managers.
func GetDeps() error {
	err := sh.RunV("go", "get", "-u", "golang.org/x/vgo")
	if err != nil {
		return err
	}
	return sh.RunV("vgo", "vendor")
}

// Clean files and folders used for test data or binaries
func Clean() error {
	err := sh.Rm("golog")
	if err != nil {
		return err
	}
	err = sh.Rm("coverage.txt")
	if err != nil {
		return err
	}
	return sh.Rm("vendor")

}

// Run tests
func Test() error {
	args := []string{"test", "-race", "-coverprofile=coverage.txt", "-covermode=atomic", "./..."}
	if mg.Verbose() {
		args = append(args, "-v")
	}
	return sh.RunV("vgo", args...)
}

// Generates a new release. Expects the TAG environment variable to be set,
// which will create a new tag with that name.
func Release() (err error) {
	if os.Getenv("TAG") == "" {
		return errors.New("MSG and TAG environment variables are required")
	}
	if err := sh.RunV("git", "tag", "-a", "$TAG"); err != nil {
		return err
	}
	if err := sh.RunV("git", "push", "origin", "$TAG"); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err = sh.RunV("git", "tag", "--delete", "$TAG")
			err = sh.RunV("git", "push", "--delete", "origin", "$TAG")
		}
	}()
	return sh.RunV("goreleaser")
}

// Clean, Build, Tests
func All() {
	mg.SerialDeps(Clean, Build, Test)
}
