package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterFileUpdate() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForFile(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForFileUpdate()),
		func(duktapeContext *duktape.Context) int {
			from := duktapeContext.RequireString(0)
			to := from
			if FileOrigin(from) {
				panic(fmt.Sprintf("not update a origin file(%s)", from))
			}
			idsJson := ""
			if duktapeContext.GetTop() > 1 {
				idsJson = duktapeContext.RequireString(1)
			}
			FileInsert(from, to, idsJson, duktapeContext)
			return 1
		})
}
