package main

import (
	"speedup/document"
	fs "speedup/filesystem"
	idx "speedup/wordprocess/indexwriter"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func main() {

	/**
	s := make([][]uint, 0)
	s = append(s, make([]uint,2))

	bolB, _ := json.Marshal(s)
	fmt.Println(string(bolB))

	if true {
		return
	}
	**/

	doc := new(document.Document).CreateDocument(1)
	doc.AddField("nome", "thiago. luiz çao rodrigues, thiago carroca")
	doc.AddField("taliba", "thiago. luiz çao rodrigues, thiago")
	doc.AddField("t2", "thiago. luiz çao rodrigues, thiago")
	doc.AddField("idade", 25)
	doc.AddField("email", "bobboyms@gmail.com")

	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem()
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	IndexWriter.IndexDocument(doc)

	//criar uma função chamada indexDocument

}
