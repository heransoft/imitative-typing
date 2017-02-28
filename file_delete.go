package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
	"os"
)

func RegisterFileDelete() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForFile(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForFileDelete()),
		func(duktapeContext *duktape.Context) int {
			file := duktapeContext.RequireString(0)
			if FileOrigin(file) {
				panic(fmt.Sprintf("not delete origin file(%s)", file))
			}
			err := os.Remove(file)
			if err != nil {
				panic(err)
			}
			return 1
		})
}
