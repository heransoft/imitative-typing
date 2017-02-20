package imitative_typing

import (
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterDirectorySelect() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForDirectory()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForDirectorySelect(),
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
