package imitative_typing

import (
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineDelete() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForLine()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForLineDelete(),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(duktapeContext.RequireInt(0))
			imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Delete[lineIndex] = true
			return 1
		})
}
