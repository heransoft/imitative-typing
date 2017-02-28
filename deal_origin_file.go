package imitative_typing

func DealOriginFile(filename string) {
	javaScriptContext.duktapeContext.EvalFile(filename)
}
