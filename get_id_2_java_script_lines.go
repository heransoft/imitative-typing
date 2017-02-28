package imitative_typing

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"strings"
)

func GetID2JavaScriptLines(filename string) (id2JavaScriptLines map[string]*JavaScriptLines) {
	id2JavaScriptLines = make(map[string]*JavaScriptLines)
	allLine := ReadFileLine(filename)
	allLineLength := len(allLine)
	annotationSymbol := GetAnnotationSymbol(filename)
	annotationSymbolLength := len(annotationSymbol)
	javaScriptStartSymbol := annotationSymbol + imitativeTypingConfig.GetJavaScriptStartSymbol()
	javaScriptStartSymbolLength := len(javaScriptStartSymbol)
	javaScriptEndSymbol := annotationSymbol + imitativeTypingConfig.GetJavaScriptEndSymbol()
	for i := 0; i < allLineLength; i++ {
		line := strings.TrimSpace(allLine[i])
		if strings.Index(line, javaScriptStartSymbol) == 0 {
			javaScriptLines := &JavaScriptLines{
				BaseLineNumber: proto.Uint32(uint32(i)),
				Order:          proto.Int32(int32(i)),
			}
			if javaScriptStartSymbolLength != len(line) {
				//unmarshal json to id and order
				jsonStr := line[javaScriptStartSymbolLength:]
				data := make(map[string]interface{}, 0)
				err := json.Unmarshal([]byte(jsonStr), data)
				if err != nil {
					data = make(map[string]interface{}, 0)
				}
				if id, exist := data["i"]; exist {
					id2JavaScriptLines[fmt.Sprintf("%v", id)] = javaScriptLines
				} else {
					id2JavaScriptLines[fmt.Sprint(i)] = javaScriptLines
				}

				if order, exist := data["o"]; exist {
					value, err := strconv.ParseInt(fmt.Sprintf("%v", order), 10, 0)
					if err == nil {
						javaScriptLines.Order = proto.Int32(int32(value))
					}
				}
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
