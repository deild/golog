// +build mage

// nolint
package main

import (
	"errors"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// A build step
func Build() error {
	mg.Deps(GetDeps)
	return sh.RunV("go", "build", "-o", "golog", ".")
}

// Manage your deps, or running package managers.
func GetDeps() error {
	return sh.RunV("dep", "ensure")
}

// Clean files and folders used for test data or binaries.
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

// Run tests.
func Test() error {
	args := []string{"test", "-race", "-coverprofile=coverage.txt", "-covermode=atomic", "./..."}
	if mg.Verbose() {
		args = append(args, "-v")
	}
	return sh.RunV("go", args...)
}

// Generates a new release. Expects the TAG environment variable to be set,
// which will create a new tag with that name.
func Release() (err error) {
	if os.Getenv("TAG") == "" {
		return errors.New("MSG and TAG environment variables are required")
	}
	if err := sh.RunV("git", "tag", "-a", "$TAG", "-m", "$TAG"); err != nil {
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
	return sh.RunV("goreleaser", "--rm-dist")
}

// Generates a new snapshot.
func Snapshot() error {
	return sh.RunV("goreleaser", "--rm-dist", "--snapshot")
}

// Clean, Build, Tests.
func All() {
	mg.SerialDeps(Clean, Build, Test)
}
