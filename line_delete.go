package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineDelete() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForLine(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForLineDelete()),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(duktapeContext.RequireInt(0))
			imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Delete[lineIndex] = true
			return 1
		})
}
