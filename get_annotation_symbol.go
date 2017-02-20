package imitative_typing

import (
	"path"
)

func GetAnnotationSymbol(filename string) string {
	fileExtension := path.Ext(filename)
	for _, v := range imitativeTypingConfig.FileExtension2AnnotationSymbol {
		if v.GetFileExtension() == fileExtension {
			return v.GetAnnotationSymbol()
		}
	}
	return "//"

}
