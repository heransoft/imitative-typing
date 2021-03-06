package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterLineUpdate() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForLine(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForLineUpdate()),
		func(duktapeContext *duktape.Context) int {
			lineIndex := uint32(duktapeContext.RequireInt(0))
			updateString := duktapeContext.RequireString(1)
			imitativeTypingContext.FileContext[imitativeTypingContext.CurrentFileName].Update[lineIndex] = updateString
			return 1
		})
}
