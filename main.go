package main

import (
	"fmt"
	"github.com/draftcode/sandal/lang"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Specify a source")
		os.Exit(1)
	}
	filePath := os.Args[1]
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(filePath, err)
	}
	err, compiled := lang.CompileFile(string(body))
	if err != nil {
		log.Fatal(filePath, err)
	}
	fmt.Print(compiled)
}
