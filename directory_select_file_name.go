package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterDirectorySelectFileName() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForDirectory(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForDirectorySelectFileName()),
		func(duktapeContext *duktape.Context) int {
			if duktapeContext.GetTop() > 1 {
				fileIndex := uint32(duktapeContext.RequireInt(0))
				rootPath := duktapeContext.RequireString(1)
				if directoryContext, exist := imitativeTypingContext.DirectoryContext[rootPath]; exist {
					filename, exist := directoryContext.FileIndex2FileName[fileIndex]
					if exist {
						duktapeContext.PushString(filename)
					} else {
						duktapeContext.PushString("")
					}
				} else {
					duktapeContext.PushString("")
				}
			} else {
				fileIndex := uint32(duktapeContext.RequireInt(0))
				filename, exist := imitativeTypingContext.FileIndex2FileName[fileIndex]
				if exist {
					duktapeContext.PushString(filename)
				} else {
					duktapeContext.PushString("")
				}
			}
			return 1
		})
}
