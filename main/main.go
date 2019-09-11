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
	/**
		list := make([][]string, 0)

		arr := []string{"2", "4", "8"}
		list = append(list, arr)

		arr = []string{"3", "2", "6"}
		list = append(list, arr)

		arr = []string{"7", "6", "9"}
		list = append(list, arr)

		result := list[0]
		for i := 1; i <= len(list)-1; i++ {

			result = difference(result, list[i])

		}

		println("Total result: ", len(result))

		for _, v := range result {
			println(v)
		}

		if true {
			return
		}
	**/
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

	start := time.Now()
	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem("contas_medicas", workFolder)
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	qr := new(query.Query).CreateQuery(fileSystem)

	//qr.Find("IDADE", "49")

	qr.AddEq(&query.Equal{
		Key:   "IDADE",
		Value: "20",
	})

	qr.AddEq(&query.Equal{
		Key:   "CSEXOUSUA",
		Value: "F",
	})

	qr.AddEq(&query.Equal{
		Key:   "NQUANTCON",
		Value: "1",
	})

	qr.FilterAnd(qr)

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
