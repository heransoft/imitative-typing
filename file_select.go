package imitative_typing

import (
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterFileSelect() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForFile()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForFileSelect(),
		func(duktapeContext *duktape.Context) int {
			fileIndex := uint32(duktapeContext.RequireInt(0))
			filename, exist := imitativeTypingContext.FileIndex2FileName[fileIndex]
			if exist {
				duktapeContext.PushString(filename)
			} else {
				duktapeContext.PushString("")
			}
			return 1
		})
}
