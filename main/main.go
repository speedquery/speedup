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

	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem()
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	doc := new(document.Document).CreateDocument(1)
	doc.AddField("nome", "thiago luiz")
	doc.AddField("idade", 25)
	//doc.AddField("email", "bobboyms@gmail.com")

	IndexWriter.IndexDocument(doc)

	println("==================================")

	doc2 := new(document.Document).CreateDocument(2)
	doc2.AddField("nome", "thiago luiz")
	doc2.AddField("idade", 30)

	//doc2.AddField("endereco", "thiago. luiz çao rodrigues, thiago taliba")
	//doc2.AddField("cidade", "rolim de moura")

	IndexWriter.IndexDocument(doc2)

	//criar uma função chamada indexDocument

}
