package filesystem

/*
FileSystem tem a função de fazer o gerenciamento de todos os indices
criados
*/
type FileSystem struct {
	wordmap      *WordMap
	attributeMap *AttributeMap
}

func (fs *FileSystem) CreateFileSystem() *FileSystem {
	fs.wordmap = new(WordMap).InitWordMap()
	fs.attributeMap = new(AttributeMap).IniAttributeMap()
	return fs
}

func (fs *FileSystem) GetWordMap() *WordMap {
	return fs.wordmap
}

func (fs *FileSystem) GetAttributeMap() *AttributeMap {
	return fs.attributeMap
}
