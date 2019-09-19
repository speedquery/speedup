package query

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"speedup/collection"
	"speedup/filesystem"
	"speedup/utils"
	"speedup/wordprocess/stringprocess"
	"strconv"
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
	andList    []Operators
}

func (self *Query) CreateQuery(filesystem *filesystem.FileSystem) *Query {
	self.filesystem = filesystem
	return self
}

func (self *Query) FilterOr(query *OR) []string {

	listEq := query.GetList()

	var result []string
	for _, value := range listEq {

		result = self.FilterAnd(value)

		if len(result) > 0 {
			break
		}

	}

	return result

}

func (self *Query) AddAnd(query Operators) *Query {

	if self.andList == nil {
		self.andList = make([]Operators, 0)
	}

	self.andList = append(self.andList, query)

	return self
}

func (self *Query) FilterAnd(query Operators) []string {

	var wg sync.WaitGroup

	list := make([][]string, 0)

	for _, eq := range query.GetList() {

		key := eq.Key
		value := eq.Value

		wg.Add(1)
		go func(key, value string, onClose func()) {

			defer onClose()

			switch query.(type) {
			case *EQ:

				result := self.FindEQ(key, value)

				if len(result) > 0 {
					list = append(list, result)
				}

			case *NotEQ:

				result := self.FindNotEQ(key, value)

				if len(result) > 0 {
					list = append(list, result)
				}

			}

		}(key, value, func() { wg.Done() })
	}

	wg.Wait()

	if len(list) > 0 {
		result := list[0]
		for i := 1; i <= len(list)-1; i++ {

			result = difference(result, list[i])

		}

		return result
	} else {
		return nil
	}

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

func (self *Query) FindEQ(key, value string) []string {

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		panic("ATRIBUTO NAO ENCONTRADO")
	}

	words := strings.Split(value, " ")

	wordGroup := make([]string, 0) //list.New()

	for _, word := range words {
		newWord := stringprocess.ProcessWord(word)
		idword := self.filesystem.GetWordMap().AddWord(newWord)
		wordGroup = append(wordGroup, fmt.Sprint(*idword))
	}

	idWordGroup := self.filesystem.GetWordGroupMap().GetAWordGroup(strings.Join(wordGroup, ""))

	if idWordGroup == nil {
		return nil
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

	return result

}

func (self *Query) FindNotEQ(key, value string) []string {

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

	strWordGroup := strings.Join(wordGroup, "")

	idWordGroup := self.filesystem.GetWordGroupMap().GetAWordGroup(strWordGroup)

	if idWordGroup == nil {
		return nil
	}

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedindex" + GetBar() + fmt.Sprintf("%v", *idWordGroup) + ".txt"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ignoredDocuments := make([]uint, 0)
	for scanner.Scan() {

		rs, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		ignoredDocuments = append(ignoredDocuments, uint(rs))

	}

	docsResult := self.filesystem.GetDocumentGroupWord().GetMapIgnoreKeys(ignoredDocuments)

	for key, _ := range docsResult {
		rs := fmt.Sprintf("%v", key)
		result = append(result, rs)
	}

	return result

}

func (self *Query) FindGT(key, value string) []string {

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		panic("ATRIBUTO NAO ENCONTRADO")
	}

	listword := self.filesystem.GetWordDocument().GetList()

	//wordGroup := make([]string, 0) //list.New()

	for _, idWord := range listword {

		word := self.filesystem.GetWordMap().GetValue(idWord)

		if !utils.IsNumber(*word) {
			continue
		}

		referenceValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			panic(err)
		}

		wordValue, err := strconv.ParseFloat(*word, 64)

		if err != nil {
			panic(err)
		}

		if wordValue <= referenceValue {
			continue
		}

		//println("Word", *word)

		path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		setDocuments := new(collection.StrSet).NewSet()
		for scanner.Scan() {
			rs := scanner.Text()
			setDocuments.Add(rs)

			//result = append(result, rs)
			//println(scanner.Text())
		}

		for k, _ := range setDocuments.GetSet() {
			result = append(result, k)
		}

	}

	return result
}

func (self *Query) FindGE(key, value string) []string {

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		panic("ATRIBUTO NAO ENCONTRADO")
	}

	listword := self.filesystem.GetWordDocument().GetList()

	//wordGroup := make([]string, 0) //list.New()

	for _, idWord := range listword {

		word := self.filesystem.GetWordMap().GetValue(idWord)

		if !utils.IsNumber(*word) {
			continue
		}

		referenceValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			panic(err)
		}

		wordValue, err := strconv.ParseFloat(*word, 64)

		if err != nil {
			panic(err)
		}

		if wordValue >= referenceValue {
			continue
		}

		path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		setDocuments := new(collection.StrSet).NewSet()
		for scanner.Scan() {
			rs := scanner.Text()
			setDocuments.Add(rs)

			//result = append(result, rs)
			//println(scanner.Text())
		}

		for k, _ := range setDocuments.GetSet() {
			result = append(result, k)
		}

	}

	return result
}
