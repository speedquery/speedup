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

	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	workFolder := userHomeDir + "\\.speedquery"

	if _, err := os.Stat(workFolder); os.IsNotExist(err) {
		os.Mkdir(workFolder, 0777)
	}

	println("Global Path:", workFolder)

	start := time.Now()
	//cria o sistema de arquivos que vai gerenciar os indices
	fileSystem := new(fs.FileSystem).CreateFileSystem("contas_medicas", workFolder)
	IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	//file, err := os.Open("C:/Users/Thago Rodrigues/go/src/speedup/dados.txt")
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
	for scanner.Scan() { // internally, it advances token based on sperator
		//fmt.Println(scanner.Text())  // token in unicode-char
		//fmt.Println(scanner.Bytes()) // token in bytes
		id++
		doc := new(document.Document).CreateDocument(id)

		flat, _ := fs.FlattenString(scanner.Text(), "", fs.DotStyle)
		doc.ToMap(flat)
		wg.Add(1)
		start := time.Now()
		IndexWriter.IndexDocument(doc, func() { wg.Done() })

		if i == 10000 {
			log.Printf("Binomial took %s", time.Since(start))
			//println("Valor de i", id)
			i = 0
		} else {
			i++
		}

		//doc.DeleteMemory()
		//println(doc.ToJson())
	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	println("Total de registro", id)

	fileSystem.SerealizeAttributeMap()
	fileSystem.SerealizeWordMap()
	fileSystem.SerealizeWordGroupMap()
	fileSystem.SerealizeAttributeGroupWord()

	//	for {
	//		time.Sleep(time.Second)
	//	}

}
