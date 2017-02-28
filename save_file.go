package imitative_typing

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func SaveFile(from, to string) {
	fileContext, exist := imitativeTypingContext.FileContext[from]
	if !exist {
		panic(fmt.Sprintf("save file(%s) failed, not exist in imitativeTypingContext.FileContext(%v)", from, imitativeTypingContext.FileContext))
	}
	file, err := os.Create(to)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	AllLineCount := uint32(len(fileContext.File.AllLine))
	for i, line := range fileContext.File.AllLine {
		lineNumber := uint32(i)
		insert, insertExist := fileContext.Insert[lineNumber]
		delete, deleteExist := fileContext.Delete[lineNumber]
		update, updateExist := fileContext.Update[lineNumber]
		if deleteExist && delete && updateExist {
			panic("deleteExist && delete && updateExist")
		}
		if insertExist {
			for _, insertLine := range insert {
				_, err = io.WriteString(file, insertLine)
				if err != nil {
					panic(err)
				}
			}
		}
		if updateExist {
			_, err = io.WriteString(file, update)
			if err != nil {
				panic(err)
			}
		} else if (deleteExist && delete) == false {
			_, err = io.WriteString(file, line)
			if err != nil {
				panic(err)
			}
		}
	}
	lineNumbers := make([]uint32, 0)
	for lineNumber := range fileContext.Insert {
		if lineNumber >= AllLineCount {
			lineNumbers = append(lineNumbers, lineNumber)
		}
	}
	sort.Sort(LineNumbersForSort(lineNumbers))
	for _, lineNumber := range lineNumbers {
		insert, insertExist := fileContext.Insert[lineNumber]
		if insertExist {
			for _, insertLine := range insert {
				_, err = io.WriteString(file, insertLine)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

type LineNumbersForSort []uint32

func (m LineNumbersForSort) Len() int { return len(m) }
func (m LineNumbersForSort) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m LineNumbersForSort) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
