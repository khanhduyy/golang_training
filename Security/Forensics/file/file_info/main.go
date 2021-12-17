package main

import (
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File name:", fileInfo.Name())
	log.Println("Size in bytes:", fileInfo.Size())
	log.Println("Permissions:", fileInfo.Mode())
	log.Println("Last modified:", fileInfo.ModTime())
	log.Println("Is Directory: ", fileInfo.IsDir())
	log.Printf("System interface type: %T\n", fileInfo.Sys())
	log.Printf("System info: %+v\n\n", fileInfo.Sys())
}
