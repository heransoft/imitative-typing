package imitative_typing

type FileContext struct {
	File       *File
	Insert     map[uint32][]string
	Delete     map[uint32]bool
	Update     map[uint32]string
	Origin     map[uint32]string
	JavaScript map[string]string
}

type DirectoryContext struct {
	FileIndex2FileName map[uint32]string
	FileName2FileIndex map[string]uint32
}

type ImitativeTypingContext struct {
	DirectoryContext   map[string]*DirectoryContext
	FileIndex2FileName map[uint32]string
	FileName2FileIndex map[string]uint32
	FileContext        map[string]*FileContext
	JavaScript         map[string]string
	CurrentFileName    string
}
