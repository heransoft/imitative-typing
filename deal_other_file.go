package imitative_typing

import "fmt"

func DealOtherFile(filename string) {
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("f.u(\"%s\")", filename))
}
