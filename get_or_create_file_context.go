package imitative_typing

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func GetOrCreateFileContext(filename string) *FileContext {
	imitativeTypingContext.CurrentFileName = filename
	fileContext, exist := imitativeTypingContext.FileContext[filename]
	if !exist {
		fileContext = &FileContext{
			File: &File{
				AllLine: ReadFileLine(filename),
				Name:    proto.String(filename),
			},
			Insert: make(map[uint32][]string),
			Delete: make(map[uint32]bool),
			Update: make(map[uint32]string),
		}
		fileContext.Origin = GetOriginLines(fileContext.File)
		fileContext.JavaScript = GetID2JavaScript(fileContext.File)
		for id, javaScript := range fileContext.JavaScript {
			if _, exist := imitativeTypingContext.JavaScript[id]; exist {
				panic(fmt.Sprintf("id(%s) duplication", id))
			}
			imitativeTypingContext.JavaScript[id] = javaScript
		}
		imitativeTypingContext.FileContext[filename] = fileContext
	}
	return fileContext
}
