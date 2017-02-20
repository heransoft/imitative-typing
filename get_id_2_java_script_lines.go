package imitative_typing

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
)

func GetID2JavaScriptLines(filename string) (id2JavaScriptLines map[string]*JavaScriptLines) {
	id2JavaScriptLines = make(map[string]*JavaScriptLines)
	allLine := ReadFileLine(filename)
	allLineLength := len(allLine)
	annotationSymbol := GetAnnotationSymbol(filename)
	annotationSymbolLength := len(annotationSymbol)
	javaScriptStartSymbol := annotationSymbol + imitativeTypingConfig.GetJavaScriptStartSymbol()
	javaScriptStartSymbol4ID := javaScriptStartSymbol + "id:"
	javaScriptStartSymbol4IDLength := len(javaScriptStartSymbol4ID)
	javaScriptEndSymbol := annotationSymbol + imitativeTypingConfig.GetJavaScriptEndSymbol()
	for i := 0; i < allLineLength; i++ {
		line := strings.TrimSpace(allLine[i])
		if strings.Index(line, javaScriptStartSymbol) == 0 {
			javaScriptLines := &JavaScriptLines{
				BaseLineNumber: proto.Uint32(uint32(i)),
			}
			if strings.Index(line, javaScriptStartSymbol4ID) == 0 {
				id := line[javaScriptStartSymbol4IDLength:]
				if _, exist := id2JavaScriptLines[id]; exist {
					panic(fmt.Sprintf("id(%s) duplication", id))
				}
				id2JavaScriptLines[id] = javaScriptLines
			} else {
				id2JavaScriptLines[fmt.Sprint(i)] = javaScriptLines
			}
			for {
				i++
				line = strings.TrimSpace(allLine[i])
				if strings.Index(line, annotationSymbol) == 0 {
					if strings.Index(line, javaScriptEndSymbol) == 0 {
						break
					}
					javaScriptLines.Lines = append(javaScriptLines.Lines, line[annotationSymbolLength:])
				} else {
					break
				}
			}
		}
	}
	return
}
