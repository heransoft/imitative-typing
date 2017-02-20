package imitative_typing

import (
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineOrigin() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForLine()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForLineOrigin(),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(duktapeContext.RequireInt(0))
			_, exist := imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Origin[lineIndex]
			duktapeContext.PushBoolean(exist)
			return 1
		})
}
