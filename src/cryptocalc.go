package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	// "github.com/emicklei/hopwatch"
	"io"
	"os"
)

var Usage = func() {
	fmt.Println("Usage: cryptocalc command filepath/filename")
	fmt.Println("The command are: md5 sha1 sha256 sha512")
}

func main() {
	args := os.Args
	if args == nil || len(args) != 3 {
		Usage()
		return
	}

	File := args[2]
	infile, err := os.Open(File)
	if err == nil {
		switch args[1] {
		case "md5":
			hash := md5.New()
			io.Copy(hash, infile)
			fmt.Printf("MD5: %x\n", hash.Sum([]byte("")))
		case "sha1":
			hash := sha1.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA1: %x\n", hash.Sum([]byte("")))
		case "sha256":
			hash := sha256.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA256: %x\n", hash.Sum([]byte("")))
		case "sha512":
			hash := sha512.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA512: %x\n", hash.Sum([]byte("")))
		default:
			Usage()
		}
	}
	return
}
