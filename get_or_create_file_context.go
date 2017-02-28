package imitative_typing

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func GetOrCreateFileContext(filename string) *FileContext {
	imitativeTypingContext.CurrentFileName = filename
	fileContext, exist := imitativeTypingContext.FileContext[filename]
	if filename == "" { //for file insert("","1.txt","[\"xx\"]")
		fileContext = &FileContext{
			File: &File{
				AllLine: []string{},
				Name:    proto.String(filename),
			},
			Insert: make(map[uint32][]string),
			Delete: make(map[uint32]bool),
			Update: make(map[uint32]string),
		}
		fileContext.Origin = make(map[uint32]string)
		fileContext.JavaScript = make(map[string]string)
		imitativeTypingContext.FileContext[filename] = fileContext
	} else if !exist {
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
		fileContext.JavaScript, fileContext.JavaScriptOrder = GetID2JavaScript(fileContext.File)
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
