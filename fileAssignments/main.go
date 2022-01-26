package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
 f, err := os.Open(os.Args[1])
 if err != nil {
	 fmt.Println("Error:", err)
	 os.Exit(1)
 }
 io.Copy(os.Stdout, f)
}


// func Open(name string, flag int, perm os.FileMode) (*File, error)