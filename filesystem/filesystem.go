package filesystem

/*
FileSystem tem a função de fazer o gerenciamento de todos os indices
criados
*/
type FileSystem struct {
	wordmap            *WordMap
	attributeMap       *AttributeMap
	attributeWord      *AttributeWord
	wordGroupMap       *WordGroupMap
	attributeGroupWord *AttributeGroupWord
	groupWordDocument  *GroupWordDocument
}

func (fs *FileSystem) CreateFileSystem() *FileSystem {
	fs.wordmap = new(WordMap).InitWordMap()
	fs.attributeMap = new(AttributeMap).IniAttributeMap()
	fs.attributeWord = new(AttributeWord).InitAttributeWord()
	fs.wordGroupMap = new(WordGroupMap).IniWordGroupMap()
	fs.attributeGroupWord = new(AttributeGroupWord).InitAttributeGroupWord()
	fs.groupWordDocument = new(GroupWordDocument).InitGroupWordDocument()
	return fs
}

func (fs *FileSystem) GetWordMap() *WordMap {
	return fs.wordmap
}

func (fs *FileSystem) GetAttributeMap() *AttributeMap {
	return fs.attributeMap
}

func (fs *FileSystem) GetAttributeWord() *AttributeWord {
	return fs.attributeWord
}

func (fs *FileSystem) GetWordGroupMap() *WordGroupMap {
	return fs.wordGroupMap
}

func (fs *FileSystem) GetAttributeGroupWord() *AttributeGroupWord {
	return fs.attributeGroupWord
}

func (fs *FileSystem) GetGroupWordDocument() *GroupWordDocument {
	return fs.groupWordDocument
}
