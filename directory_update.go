package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterDirectoryUpdate() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForDirectory(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForDirectoryUpdate()),
		func(duktapeContext *duktape.Context) int {
			DirectoryInsert(duktapeContext)
			DirectoryUpdate()
			return 1
		})
}

func DirectoryUpdate() {
	imitativeTypingContext.FileIndex2FileName = make(map[uint32]string)
	imitativeTypingContext.FileName2FileIndex = make(map[string]uint32)
	for _, directoryContext := range imitativeTypingContext.DirectoryContext {
		for _, path := range directoryContext.FileIndex2FileName {
			if _, exist := imitativeTypingContext.FileName2FileIndex[path]; exist == false {
				index := uint32(len(imitativeTypingContext.FileIndex2FileName))
				imitativeTypingContext.FileIndex2FileName[index] = path
				imitativeTypingContext.FileName2FileIndex[path] = index
			}
		}
	}
}
