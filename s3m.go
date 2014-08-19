package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		usage(nil)
	}

	command := os.Args[1]
	var err error
	switch command {
	case "clone":
		err = clone(os.Args[2:])
	case "pull":
		err = pull()
	case "push":
		err = push()
	case "diff":
		err = diff()
	default:
		err = fmt.Errorf("Invalid command.")
	}

	if err != nil {
		usage(err)
	}
}

func usage(err error) {
	if err != nil {
		fmt.Printf("%s: ERROR: %s\n", filepath.Base(os.Args[0]), err)
	}
	fmt.Printf("usage: %s {command} [arguments]\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func clone(args []string) error {

	var s3Path string
	var localPath string

	if len(args) < 1 {
		return fmt.Errorf("Clone command requires an s3 bucket value")
	}
	s3Path = args[0]

	localPath, _ = filepath.Abs(filepath.Base(s3Path))
	if len(args) > 1 {
		localPath, _ = filepath.Abs(args[1])
	}

	fmt.Printf("clone s3://%s to %s\n", s3Path, localPath)

	return nil
}

func pull() error {
	fmt.Println("PULL")
	return nil
}

func push() error {
	fmt.Println("PUSH")
	return nil
}

func diff() error {
	fmt.Println("DIFF")
	return nil
}
