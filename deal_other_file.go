package imitative_typing

import "fmt"

func DealOtherFile(filename string) {
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("%s.%s(\"%s\")",
		imitativeTypingConfig.GetJavaScriptTableNameForFile(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForFileUpdate(),
		filename))
}
