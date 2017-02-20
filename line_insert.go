package imitative_typing

import (
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineInsert() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForLine()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForLineInsert(),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(duktapeContext.RequireInt(0))
			insertString := duktapeContext.RequireString(1)
			insertMap := imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Insert[lineIndex]
			insertMap = append(insertMap, insertString)
			imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Insert[lineIndex] = insertMap
			return 1
		})
}
