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

	doc := new(document.Document).CreateDocument(1)
	doc.AddField("nome", "thiago. luiz çao rodrigues")
	doc.AddField("idade", 25)
	doc.AddField("email", "bobboyms@gmail.com")

	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem()
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	IndexWriter.IndexDocument(doc)

	//criar uma função chamada indexDocument

}
