package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
	"os"
	"path/filepath"
	"regexp"
)

func RegisterDirectoryInsert() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForDirectory(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForDirectoryInsert()),
		DirectoryInsert)
}

func DirectoryInsert(duktapeContext *duktape.Context) int {
	rootPath := duktapeContext.RequireString(0)
	imitativeTypingContext.DirectoryContext[rootPath] = &DirectoryContext{
		FileIndex2FileName: make(map[uint32]string),
		FileName2FileIndex: make(map[string]uint32),
	}
	reg := regexp.MustCompile(".*")
	if duktapeContext.GetTop() > 1 {
		reg = regexp.MustCompile(duktapeContext.RequireString(1))
	}
	err := filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if reg.MatchString(path) == false {
			return nil
		}
		{
			if _, exist := imitativeTypingContext.FileName2FileIndex[path]; exist == false {
				index := uint32(len(imitativeTypingContext.FileIndex2FileName))
				imitativeTypingContext.FileIndex2FileName[index] = path
				imitativeTypingContext.FileName2FileIndex[path] = index
			}
		}
		{
			index := uint32(len(imitativeTypingContext.DirectoryContext[rootPath].FileIndex2FileName))
			imitativeTypingContext.DirectoryContext[rootPath].FileIndex2FileName[index] = path
			imitativeTypingContext.DirectoryContext[rootPath].FileName2FileIndex[path] = index
		}
		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("filepath.Walk() returned %v\n", err))
	}
	return 1
}
