package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"speedup/document"
	fs "speedup/filesystem"
	"speedup/query/newquery"
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

	file, err := os.Open("speedup/dados.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//buf := make([]byte, 0, 1024*1024)
	//scanner.Buffer(buf, 10*1024*1024)

	//var id uint
	id := uint(0)

	var wg sync.WaitGroup
	var i uint = 0

	start := time.Now()
	for scanner.Scan() {

		id++
		doc := new(document.Document).CreateDocument(id)

		//println(scanner.Text())

		flat, _ := fs.FlattenString(scanner.Text(), "", fs.DotStyle)
		doc.ToMap(flat)
		wg.Add(1)
		start := time.Now()

		IndexWriter.IndexDocument(doc, true)
		doc = doc.DeleteMemoryDocument()

		if i == 10000 {
			log.Printf("Binomial took %s", time.Since(start))
			i = 0
		} else {
			i++
		}
	}

	//wg.Wait()

	println("---- Concluido ----")
	log.Printf("Binomial took %s", time.Since(start))
	time.Sleep(time.Minute * 2)

	if true {
		return
	}

	for {
		time.Sleep(time.Minute)
	}
}

func main() {

	TesteIndexacaoTeste()

	if true {
		return
	}

	workFolder := utils.InitializeWorkFolder()
	fileSystem := new(fs.FileSystem).CreateFileSystem("contas_medicas", workFolder)
	//IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	println("iniciou a query")

	start := time.Now()

	qr := new(newquery.QUERY).Create(fileSystem)

	Andgroup := new(newquery.GROUP).AddOperator(new(newquery.GEE).Add(&newquery.Map{
		Key:   "VALORPAGO",
		Value: "1",
	}))

	// .AddOperator(new(newquery.EQ).Add(&newquery.Map{
	// 	Key:   "nome",
	// 	Value: "isabella rodrigues",
	// }))

	/**
	ORgroup := new(newquery.GROUP).AddOperator(new(newquery.EQ).Add(&newquery.Map{
		Key:   "idade",
		Value: "30",
	}))
	**/

	rs := qr.Add(new(newquery.AND).AddGroup(Andgroup)).GetList()

	//.Add(new(newquery.OR).AddGroup(ORgroup)).GetList()

	println(rs == nil)

	log.Printf("Binomial took %s", time.Since(start))
	println("Total:", len(rs))

	// for _, r := range rs {
	// 	println(r)
	// }

	/**
	qr := new(query.Query).CreateQuery(fileSystem)

	rs := qr.Add(new(query.EQ).AddEQ(&query.Map{
		Key:   "idade",
		Value: "300",
	}).AddEQ(&query.Map{
		Key:   "nome",
		Value: "jose luiz",
	})).AddOR(new(query.OR).AddOR(new(query.EQ).AddEQ(&query.Map{
		Key:   "idade",
		Value: "30",
	}))).GetList()

	***/

	/**
		)
	***/

	// rs := qr.FindAttEQ("idade", "300")
	//rs := qr.FindIndexEQ("30")

	/**
		qq := new(query.NotEQ).AddEQ(&query.Map{
			Key:   "idade",
			Value: "30",
		})

		//rs := qr.FilterAnd(qq)
	**/

	//for _, v := range rs {
	//	println(v)
	//}

	//cria o sistema de arquivos que vai gerenciar os indices

	//for {
	//	time.Sleep(time.Second)
	//}

}
