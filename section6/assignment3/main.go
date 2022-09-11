package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
    fileName := os.Args[len(os.Args)-1]
    fmt.Println(fileName)
    byt, _ := os.OpenFile(fileName, os.O_RDONLY, 0755)
    io.Copy(os.Stdout, byt)

}