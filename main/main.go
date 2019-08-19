package main

import (
	"bufio"
	"log"
	"os"
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

	file, err := os.Open("speedup/teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var id uint
	id = 1
	for scanner.Scan() { // internally, it advances token based on sperator
		//fmt.Println(scanner.Text())  // token in unicode-char
		//fmt.Println(scanner.Bytes()) // token in bytes

		doc := new(document.Document).CreateDocument(id)

		flat, _ := fs.FlattenString(scanner.Text(), "", fs.DotStyle)
		doc.ToMap(flat)
		//println(doc.ToJson())
		IndexWriter.IndexDocument(doc)

	}

}
