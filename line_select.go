package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineSelect() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForLine(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForLineSelect()),
		func(duktapeContext *duktape.Context) int {
			lineIndex := int32(duktapeContext.RequireInt(0))
			currentFileLineContent := imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].File.AllLine[lineIndex]
			duktapeContext.PushString(currentFileLineContent)
			return 1
		})
}
