package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
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

	file, err := os.Open("speedup/teste.txt")
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

	//for i := 1; i < 10000; i++ {
	//	IndexWriter.DeleteDocument(uint(i))
	//}

	//dx := new(document.Document).CreateDocument(1)
	//dx.ToMap(`{"nome":"tatiane rodrigues", "idade":30}`)
	//IndexWriter.UpdateDocument(dx)

	//wg.Wait()

	//elapsed := time.Since(start)
	//log.Printf("Binomial took %s", elapsed)

	if true {

		for scanner.Scan() { // internally, it advances token based on sperator
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
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	println("Total de registro", id)

	for {
		time.Sleep(time.Second)
	}

}
