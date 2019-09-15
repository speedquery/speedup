package main

import (
<<<<<<< HEAD
	"bufio"
	"log"
	"os"
	"runtime"
=======
	"fmt"
	"reflect"
>>>>>>> 4904f70da0c7e5c04f9c4351f39631fe978b1cca
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

	if id != 5 {
		panic("Quantidade de registros deve ser 5")
	}

	time.Sleep(time.Minute)
}

func main() {

	workFolder := utils.InitializeWorkFolder()

	fileSystem := new(fs.FileSystem).CreateFileSystem("teste", workFolder)
	//IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	qr := new(query.Query).CreateQuery(fileSystem)

	qr.FindNotEQ("nome", "thiago luiz")

	if true {
		return
	}

	//cria o sistema de arquivos que vai gerenciar os indices

<<<<<<< HEAD
	//for {
	//	time.Sleep(time.Second)
	//}
=======
	//criar uma função chamada indexDocument
	doc = new(document.Document).CreateDocument(1)
	doc.AddField("nome", "jose taliba. luiz çao rodrigues")
	doc.AddField("idade", 54)
	doc.AddField("email", "bobboyms@gmail.com")

	IndexWriter.IndexDocument(doc)

	a := make([]int, 0)
	b := make([]int, 0)
	fmt.Println(reflect.DeepEqual(a, b))
>>>>>>> 4904f70da0c7e5c04f9c4351f39631fe978b1cca

}
