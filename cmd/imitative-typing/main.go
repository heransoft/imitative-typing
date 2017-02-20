package main

import (
	"fmt"
	"github.com/heransoft/imitative-typing"
	"os"
)

func main() {
	argsCount := len(os.Args)
	if argsCount != 2 {
		panic(fmt.Sprintf("arguments has one and only one argument,but now has %d arguments", argsCount-1))
	}
	filename := os.Args[1]
	if imitative_typing.FileOrigin(filename) {
		imitative_typing.DealOriginFile(filename)
	} else {
		imitative_typing.DealOtherFile(filename)
	}
}
