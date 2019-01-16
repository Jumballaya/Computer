package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jumballaya/computer/software/assembler"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error has occured: %q\n", err)
		return
	}

	assembler.Compile(string(file))
}
