package imitative_typing

import (
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

type JavaScriptContext struct {
	duktapeContext *duktape.Context
}

func NewJavaScriptContext() *JavaScriptContext {
	javaScriptContext := &JavaScriptContext{
		duktapeContext: duktape.New(),
	}
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("var %s = {}", imitativeTypingConfig.GetJavaScriptTableNameForGlobal()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("var %s = {}", imitativeTypingConfig.GetJavaScriptTableNameForDirectory()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("var %s = {}", imitativeTypingConfig.GetJavaScriptTableNameForFile()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("var %s = {}", imitativeTypingConfig.GetJavaScriptTableNameForLine()))
	javaScriptContext.duktapeContext.PevalString(`var console = {log:print,warn:print,error:print,info:print}`)
	javaScriptContext.PushGoFunction("Duktape.modSearch", func(c *duktape.Context) int {
		name := c.RequireString(0)
		javaScriptContext.duktapeContext.PushStringFile(fmt.Sprintf("%s.js", name))
		return 1
	})
	return javaScriptContext
}

func (javaScriptContext *JavaScriptContext) PushGoFunction(name string, function func(c *duktape.Context) int) {
	javaScriptContext.duktapeContext.EvalString(fmt.Sprintf("(function(func){%s=func;})", name))
	javaScriptContext.duktapeContext.PushGoFunction(function)
	javaScriptContext.duktapeContext.Call(1)
	javaScriptContext.duktapeContext.Pop()
}
