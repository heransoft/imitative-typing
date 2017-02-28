package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineInsert() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForLine(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForLineInsert()),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(0)
			insertString := ""
			if duktapeContext.GetTop() > 1 {
				lineIndex = uint32(duktapeContext.RequireInt(0))
				insertString = duktapeContext.RequireString(1)
			} else {
				insertString = duktapeContext.RequireString(0)
			}
			insertMap := imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Insert[lineIndex]
			insertMap = append(insertMap, insertString)
			imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Insert[lineIndex] = insertMap
			return 1
		})
}
