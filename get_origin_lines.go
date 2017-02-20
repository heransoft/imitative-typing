package imitative_typing

import (
	"strings"
)

func GetOriginLines(file *File) (originLines map[uint32]string) {
	originLines = make(map[uint32]string)
	allLine := file.AllLine
	allLineLength := uint32(len(allLine))
	annotationSymbol := GetAnnotationSymbol(file.GetName())
	javaScriptStartSymbol := annotationSymbol + imitativeTypingConfig.GetJavaScriptStartSymbol()
	javaScriptEndSymbol := annotationSymbol + imitativeTypingConfig.GetJavaScriptEndSymbol()
	for i := uint32(0); i < allLineLength; i++ {
		line := strings.TrimSpace(allLine[i])
		if strings.Index(line, javaScriptStartSymbol) == 0 {
			for {
				i++
				line = strings.TrimSpace(allLine[i])
				if strings.Index(line, annotationSymbol) == 0 {
					if strings.Index(line, javaScriptEndSymbol) == 0 {
						break
					}
				} else {
					break
				}
			}
		} else {
			originLines[i] = allLine[i]
		}
	}
	return
}
