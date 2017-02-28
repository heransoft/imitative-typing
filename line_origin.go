package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineOrigin() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForLine(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForLineOrigin()),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(duktapeContext.RequireInt(0))
			_, exist := imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Origin[lineIndex]
			duktapeContext.PushBoolean(exist)
			return 1
		})
}
