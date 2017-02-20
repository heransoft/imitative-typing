package imitative_typing

func GetID2JavaScript(file *File) (id2JavaScript map[string]string) {
	id2JavaScript = make(map[string]string)
	id2JavaScriptLines := GetID2JavaScriptLines(file.GetName())
	PreprocessorMacro(id2JavaScriptLines, file)
	for id, javaScriptLines := range id2JavaScriptLines {
		javaScript := ""
		for _, javaScriptLine := range javaScriptLines.Lines {
			javaScript += javaScriptLine
			//强制在后面加入换行
			javaScript += "\n"
		}
		id2JavaScript[id] = javaScript
	}
	return
}
