package main

import (
	"bufio"
	"log"
	"os"
	"speedup/document"
	fs "speedup/filesystem"
	idx "speedup/wordprocess/indexwriter"
	"sync"
	"time"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func main() {

	start := time.Now()

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
	id = 0

	var wg sync.WaitGroup

	for scanner.Scan() { // internally, it advances token based on sperator
		//fmt.Println(scanner.Text())  // token in unicode-char
		//fmt.Println(scanner.Bytes()) // token in bytes
		id++
		doc := new(document.Document).CreateDocument(id)

		flat, _ := fs.FlattenString(scanner.Text(), "", fs.DotStyle)
		doc.ToMap(flat)
		//println(doc.ToJson())
		wg.Add(1)
		IndexWriter.IndexDocument(doc, func() { wg.Done() })

	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed, id)

}
