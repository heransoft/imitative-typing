package imitative_typing

var imitativeTypingContext *ImitativeTypingContext
var imitativeTypingConfig *ImitativeTypingConfig
var javaScriptContext *JavaScriptContext

func init() {
	imitativeTypingConfig = &ImitativeTypingConfig{}
	LoadProto("imitative_typing_config.pb", imitativeTypingConfig)
	imitativeTypingContext = &ImitativeTypingContext{
		DirectoryContext:   make(map[string]*DirectoryContext),
		FileIndex2FileName: make(map[uint32]string),
		FileName2FileIndex: make(map[string]uint32),
		FileContext:        make(map[string]*FileContext),
		JavaScript:         make(map[string]string),
	}
	javaScriptContext = NewJavaScriptContext()

	RegisterDirectoryInsert()
	RegisterDirectoryDelete()
	RegisterDirectorySelect()
	RegisterDirectoryUpdate()
	RegisterFileInsert()
	RegisterFileDelete()
	RegisterFileSelect()
	RegisterFileUpdate()
	RegisterFileOrigin()
	RegisterLineInsert()
	RegisterLineDelete()
	RegisterLineSelect()
	RegisterLineUpdate()
	RegisterLineOrigin()
}
