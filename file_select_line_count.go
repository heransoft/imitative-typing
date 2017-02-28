package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterFileSelectLineCount() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForFile(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForFileSelectLineCount()),
		func(duktapeContext *duktape.Context) int {
			filename := ""
			if duktapeContext.GetTop() > 0 {
				filename = duktapeContext.RequireString(0)
			} else {
				filename = imitativeTypingContext.CurrentFileName
			}

			if filename == "" {
				duktapeContext.PushInt(0)
			} else {
				fileContext, exist := imitativeTypingContext.FileContext[filename]
				if exist {
					duktapeContext.PushInt(len(fileContext.File.AllLine))
				} else {
					duktapeContext.PushInt(0)
				}
			}
			return 1
		})
}
