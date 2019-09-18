package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"speedup/document"
	fs "speedup/filesystem"
	"speedup/query"
	"speedup/utils"
	idx "speedup/wordprocess/indexwriter"

	"sync"
	"time"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; found {
			diff = append(diff, x)
		}
	}
	return diff
}

func TesteIndexacaoTeste() {
	workFolder := utils.InitializeWorkFolder()

	println("OS:", runtime.GOOS, "/", runtime.GOARCH)
	println("GLOBAL PATH:", workFolder)

	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem("teste", workFolder)
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	file, err := os.Open("speedup/teste2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var id uint
	id = 0

	var wg sync.WaitGroup
	var i uint = 0

	for scanner.Scan() {

		id++
		doc := new(document.Document).CreateDocument(id)

		flat, _ := fs.FlattenString(scanner.Text(), "", fs.DotStyle)
		doc.ToMap(flat)
		wg.Add(1)
		start := time.Now()

		IndexWriter.IndexDocument(doc, false)
		doc = doc.DeleteMemoryDocument()

		if i == 10000 {
			log.Printf("Binomial took %s", time.Since(start))
			i = 0
		} else {
			i++
		}
	}

	//wg.Wait()

	time.Sleep(time.Minute * 2)
}

func main() {

	//TesteIndexacaoTeste()

	workFolder := utils.InitializeWorkFolder()

	fileSystem := new(fs.FileSystem).CreateFileSystem("teste", workFolder)
	//IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	qr := new(query.Query).CreateQuery(fileSystem)

	qr.FindGT("idade", "20")

	//qr.FindEQ("idade", "nil")

	//rs := qr.FilterAnd(noteq)

	println("------==============-----")

	if true {
		return
	}

	//cria o sistema de arquivos que vai gerenciar os indices

	//for {
	//	time.Sleep(time.Second)
	//}

}
