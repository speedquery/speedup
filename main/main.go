package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"speedup/document"
	fs "speedup/filesystem"
	"speedup/query"
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

func main() {

	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	var workFolder string

	if runtime.GOOS == "windows" {
		workFolder = userHomeDir + "\\.speedquery"
	} else {
		workFolder = userHomeDir + "/speedquery"
	}

	if _, err := os.Stat(workFolder); os.IsNotExist(err) {

		err := os.Mkdir(workFolder, 0777)

		if err != nil {
			panic(err)
		}
	}

	println("OS:", runtime.GOOS, "/", runtime.GOARCH)
	println("GLOBAL PATH:", workFolder)

	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem("contas_medicas", workFolder)
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	qr := new(query.Query).CreateQuery(fileSystem)

	condiction1 := new(query.EQ).AddEQ(&query.Map{
		Key:   "IDADE",
		Value: "20",
	})

	condiction2 := new(query.EQ).AddEQ(&query.Map{
		Key:   "CSEXOUSUA",
		Value: "M",
	}).AddEQ(&query.Map{
		Key:   "NQUANTCON",
		Value: "1",
	})

	condictionOR := new(query.OR)

	condictionOR.AddOR(condiction2).AddOR(condiction1)
	start := time.Now()

	//rs := qr.FilterAnd(condiction)
	//println("Resultado:", len(rs))

	rs := qr.FilterOr(condictionOR)
	println("Resultado:", len(rs))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	file, err := os.Open("speedup/teste2.txt")
	//file, err := os.Open("C:\\teste\\arquivos-json-completo.txt") //os.Open("speedup/dados.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var id uint
	id = 0

	var wg sync.WaitGroup
	var i uint = 0

	//IndexWriter.DeleteDocumentBulk(3)

	//for i := 1; i < 10000; i++ {
	//
	//}

	//dx := new(document.Document).CreateDocument(1)
	//dx.ToMap(`{"nome":"tatiane rodrigues", "idade":30}`)
	//IndexWriter.UpdateDocument(dx)

	//wg.Wait()

	//elapsed := time.Since(start)
	//log.Printf("Binomial took %s", elapsed)

	if false {

		for scanner.Scan() {
			//internally, it advances token based on sperator
			//fmt.Println(scanner.Text())  // token in unicode-char
			//fmt.Println(scanner.Bytes()) // token in bytes
			id++
			doc := new(document.Document).CreateDocument(id)

			flat, _ := fs.FlattenString(scanner.Text(), "", fs.DotStyle)
			doc.ToMap(flat)
			wg.Add(1)
			start := time.Now()

			//println(doc)
			IndexWriter.IndexDocument(doc, false)
			//IndexWriter.UpdateDocument(doc)
			doc = doc.DeleteMemoryDocument()

			if i == 10000 {
				log.Printf("Binomial took %s", time.Since(start))
				i = 0
			} else {
				i++
			}
		}
	}

	//wg.Wait()

	println("Total de registro", id)

	//for {
	//	time.Sleep(time.Second)
	//}

}
