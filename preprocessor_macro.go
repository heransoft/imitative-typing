package imitative_typing

import (
	"fmt"
	"strings"
)

func PreprocessorMacro(id2JavaScriptLines map[string]*JavaScriptLines, file *File) {
	fileEndLine := len(file.AllLine)
	for _, v := range id2JavaScriptLines {
		lines := v.Lines
		baseLineNumber := v.GetBaseLineNumber()
		startLine := baseLineNumber - 1
		endLine := baseLineNumber + uint32(len(lines)) + 2
		for i := range lines {
			lines[i] = strings.Replace(lines[i], imitativeTypingConfig.GetPreprocessorMacroForCurrentLine(), fmt.Sprint(baseLineNumber+uint32(i)+1), -1)
			lines[i] = strings.Replace(lines[i], imitativeTypingConfig.GetPreprocessorMacroForCurrentStartLine(), fmt.Sprint(startLine), -1)
			lines[i] = strings.Replace(lines[i], imitativeTypingConfig.GetPreprocessorMacroForCurrentEndLine(), fmt.Sprint(endLine), -1)
			lines[i] = strings.Replace(lines[i], imitativeTypingConfig.GetPreprocessorMacroForCurrentFileEndLine(), fmt.Sprint(fileEndLine), -1)
		}
	}
}
