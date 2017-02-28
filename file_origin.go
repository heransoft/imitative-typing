package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
	"strings"
)

func RegisterFileOrigin() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForFile(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForFileOrigin()),
		func(duktapeContext *duktape.Context) int {
			filename := duktapeContext.RequireString(0)
			duktapeContext.PushBoolean(FileOrigin(filename))
			return 1
		})
}

func FileOrigin(filename string) bool {
	suffix := ".it.js"
	if len(filename) <= len(suffix) {
		return false
	}
	if len(filename)-strings.Index(filename, suffix)-len(suffix) == 0 {
		return true
	}
	return false
}
