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
	javaScriptContext.duktapeContext.PevalString("var it = {}")
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("it.%s = {}", imitativeTypingConfig.GetJavaScriptTableNameForSystem()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("it.%s.%s = {}", imitativeTypingConfig.GetJavaScriptTableNameForSystem(), imitativeTypingConfig.GetJavaScriptTableNameForSystemFunctions()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("it.%s = {}", imitativeTypingConfig.GetJavaScriptTableNameForDirectory()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("it.%s = {}", imitativeTypingConfig.GetJavaScriptTableNameForFile()))
	javaScriptContext.duktapeContext.PevalString(fmt.Sprintf("it.%s = {}", imitativeTypingConfig.GetJavaScriptTableNameForLine()))
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

func RegisterFunctions() {
	RegisterDirectoryInsert()
	RegisterDirectoryDelete()
	RegisterDirectorySelectFileCount()
	RegisterDirectorySelectFileName()
	RegisterDirectoryUpdate()
	RegisterFileInsert()
	RegisterFileDelete()
	RegisterFileSelectLineCount()
	RegisterFileUpdate()
	RegisterFileOrigin()
	RegisterLineInsert()
	RegisterLineDelete()
	RegisterLineSelect()
	RegisterLineUpdate()
	RegisterLineOrigin()
}
