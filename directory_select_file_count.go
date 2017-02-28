package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterDirectorySelectFileCount() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForDirectory(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForDirectorySelectFileCount()),
		func(duktapeContext *duktape.Context) int {
			if duktapeContext.GetTop() > 0 {
				rootPath := duktapeContext.RequireString(0)
				if directoryContext, exist := imitativeTypingContext.DirectoryContext[rootPath]; exist {
					duktapeContext.PushInt(len(directoryContext.FileIndex2FileName))
				} else {
					duktapeContext.PushInt(0)
				}
			} else {
				duktapeContext.PushInt(len(imitativeTypingContext.FileIndex2FileName))
			}
			return 1
		})
}
