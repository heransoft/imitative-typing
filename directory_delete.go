package imitative_typing

import (
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterDirectoryDelete() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForDirectory()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForDirectoryDelete(),
		func(duktapeContext *duktape.Context) int {
			if duktapeContext.GetTop() > 0 {
				rootPath := duktapeContext.RequireString(0)
				if _, exist := imitativeTypingContext.DirectoryContext[rootPath]; exist {
					delete(imitativeTypingContext.DirectoryContext, rootPath)
					DirectoryUpdate()
					duktapeContext.PushBoolean(true)
				} else {
					duktapeContext.PushBoolean(false)
				}
			} else {
				imitativeTypingContext.DirectoryContext = make(map[string]*DirectoryContext)
				imitativeTypingContext.FileIndex2FileName = make(map[uint32]string)
				imitativeTypingContext.FileName2FileIndex = make(map[string]uint32)
				duktapeContext.PushBoolean(true)
			}
			return 1
		})
}
