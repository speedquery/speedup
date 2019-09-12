package query

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"speedup/filesystem"
	"speedup/wordprocess/stringprocess"
	"strings"
	"sync"
)

func GetBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

type Query struct {
	filesystem *filesystem.FileSystem
	eqlist     []*Equal
}

func (self *Query) CreateQuery(filesystem *filesystem.FileSystem) *Query {
	self.filesystem = filesystem
	return self
}

func (self *Query) AddEq(eq *Equal) *Query {

	if self.eqlist == nil {
		self.eqlist = make([]*Equal, 0)
	}

	self.eqlist = append(self.eqlist, eq)

	return self

}

func (self *Query) FilterAnd(query *Query) []string {

	var wg sync.WaitGroup

	//documents := new(collection.StrSet).NewSet()
	list := make([][]string, 0)
	for _, eq := range query.eqlist {

		key := eq.Key
		value := eq.Value

		wg.Add(1)
		go func(key, value string, onClose func()) {

			defer onClose()

			result := self.Find(key, value)

			//println("ass", len(result))

			if len(result) > 0 {
				list = append(list, result)
			}

		}(key, value, func() { wg.Done() })
	}

	wg.Wait()

	result := list[0]
	for i := 1; i <= len(list)-1; i++ {

		result = difference(result, list[i])

		println(len(result), i)

	}

	println("tamanho: ", len(result))
	//for _, vl := range result {
	//	println(vl)
	//}

	return result

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

func (self *Query) Find(key, value string) []string {

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		panic("ATRIBUTO NAO ENCONTRADO")
	}

	words := strings.Split(value, " ")

	wordGroup := make([]string, 0) //list.New()

	for _, word := range words {
		//wg.Add(1)

		newWord := stringprocess.ProcessWord(word)
		idword := self.filesystem.GetWordMap().AddWord(newWord)
		wordGroup = append(wordGroup, fmt.Sprint(*idword))
	}

	idWordGroup := self.filesystem.GetWordGroupMap().GetAWordGroup(strings.Join(wordGroup, ""))

	if idWordGroup == nil {
		panic("ERRO ID GROUP NAO ENCONTRADO")
	}

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedindex" + GetBar() + fmt.Sprintf("%v", *idWordGroup) + ".txt"

	file, err := os.Open(path)
	//file, err := os.Open("C:\\teste\\arquivos-json-completo.txt") //os.Open("speedup/dados.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		i++
		rs := scanner.Text()
		result = append(result, rs)
		//println(scanner.Text())
	}

	//println("total", i)
	//println(*idAttribute, *idWordGroup)

	//file, err := os.Open("speedup/teste2.txt")

	return result

}
