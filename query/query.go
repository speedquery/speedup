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

				result := self.FindIndexEQ(value)

				if len(result) > 0 {
					list = append(list, result)
				}

			case *NotEQ:

				result := self.FindIndexNotEQ(value)

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

func (self *Query) FindAttEQ(key, value string) []string {

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		return result
	}

	words := strings.Split(value, " ")
	wordGroup := make([]string, 0) //list.New()

	for _, word := range words {
		newWord := stringprocess.ProcessWord(word)
		idword := self.filesystem.GetWordMap().AddWord(newWord)
		wordGroup = append(wordGroup, fmt.Sprint(*idword))
	}

	idWordGroup := self.filesystem.GetWordGroupMap().GetAWordGroup(strings.Join(wordGroup, ""))
	groupWords, _ := self.filesystem.GetAttributeGroupWord().GetValue(idAttribute)

	if !groupWords.IsExistValue(idWordGroup) {
		return result
	}

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedindex" + GetBar() + fmt.Sprintf("%v", *idWordGroup) + ".txt"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rs := scanner.Text()
		result = append(result, rs)
	}

	return result

}

func (self *Query) FindAttNotEQ(key, value string) []string {

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		return result
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
		return result
	}

	groupWords, _ := self.filesystem.GetAttributeGroupWord().GetValue(idAttribute)

	if !groupWords.IsExistValue(idWordGroup) {
		return result
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

func (self *Query) FindIndexEQ(value string) []string {

	result := make([]string, 0)

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

func (self *Query) FindIndexNotEQ(value string) []string {

	result := make([]string, 0)

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
		return result
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

func (self *Query) FindAttGT(key, value string) []string {

	var wg sync.WaitGroup

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		return result
	}

	//words := strings.Split(value, " ")
	//wordGroup := make([]string, 0) //list.New()

	words, _ := self.filesystem.GetAttributeWord().GetValue(idAttribute)
	setDocuments := new(collection.StrSet).NewSet()

	count := 0

	for idWord, _ := range words.GetSet() {

		//println("ID WORD", *idWord, *self.filesystem.GetWordMap().GetValue(idWord))

		referenceValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			panic(err)
		}

		word := self.filesystem.GetWordMap().GetValue(idWord)

		if !utils.IsNumber(*word) {
			continue
		}

		wordValue, err := strconv.ParseFloat(*word, 64)

		if err != nil {
			panic(err)
		}

		if wordValue <= referenceValue {
			continue
		}

		count++
		if count > 300 {
			count = 0
			wg.Wait()
		}

		wg.Add(1)
		func(idWord *uint, onClose func()) {

			defer onClose()

			path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)

			println("ID WORD", *idWord)
			for scanner.Scan() {
				rs := scanner.Text()
				println("Documentos", rs)
				setDocuments.Add(rs)
			}

		}(idWord, func() { wg.Done() })

	}

	wg.Wait()

	for k, _ := range setDocuments.GetSet() {
		println("Set", k)
		result = append(result, k)
	}

	return result
}

func (self *Query) FindAttGE(key, value string) []string {

	var wg sync.WaitGroup

	result := make([]string, 0)

	idAttribute := self.filesystem.GetAttributeMap().GetAttribute(key)

	if idAttribute == nil {
		return result
	}

	//words := strings.Split(value, " ")
	//wordGroup := make([]string, 0) //list.New()

	words, _ := self.filesystem.GetAttributeWord().GetValue(idAttribute)
	setDocuments := new(collection.StrSet).NewSet()

	count := 0

	for idWord, _ := range words.GetSet() {

		//println("ID WORD", *idWord, *self.filesystem.GetWordMap().GetValue(idWord))

		referenceValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			panic(err)
		}

		word := self.filesystem.GetWordMap().GetValue(idWord)

		if !utils.IsNumber(*word) {
			continue
		}

		wordValue, err := strconv.ParseFloat(*word, 64)

		if err != nil {
			panic(err)
		}

		if wordValue >= referenceValue {
			continue
		}

		count++
		if count > 300 {
			count = 0
			wg.Wait()
		}

		wg.Add(1)
		func(idWord *uint, onClose func()) {

			defer onClose()

			path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)

			println("ID WORD", *idWord)
			for scanner.Scan() {
				rs := scanner.Text()
				println("Documentos", rs)
				setDocuments.Add(rs)
			}

		}(idWord, func() { wg.Done() })

	}

	wg.Wait()

	for k, _ := range setDocuments.GetSet() {
		println("Set", k)
		result = append(result, k)
	}

	return result
}

func (self *Query) FindIndexGT(value string) []string {

	var wg sync.WaitGroup

	result := make([]string, 0)

	listword := self.filesystem.GetWordDocument().GetNumberList()
	setDocuments := new(collection.StrSet).NewSet()
	//setWords := new(collection.Set).NewSet()

	count := 0

	for _, idWord := range listword {

		word := self.filesystem.GetWordMap().GetValue(idWord)

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

		count++
		if count > 300 {
			count = 0
			wg.Wait()
		}

		wg.Add(1)
		go func(idWord *uint, onClose func()) {

			defer onClose()

			path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				rs := scanner.Text()
				setDocuments.Add(rs)
			}

		}(idWord, func() { wg.Done() })

	}

	wg.Wait()

	for k, _ := range setDocuments.GetSet() {
		result = append(result, k)
	}

	return result
}

func (self *Query) FindIndexGE(value string) []string {

	var wg sync.WaitGroup

	result := make([]string, 0)

	listword := self.filesystem.GetWordDocument().GetNumberList()
	setDocuments := new(collection.StrSet).NewSet()
	//setWords := new(collection.Set).NewSet()

	count := 0

	for _, idWord := range listword {

		word := self.filesystem.GetWordMap().GetValue(idWord)

		//println(*word)

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

		count++
		if count > 300 {
			count = 0
			wg.Wait()
		}

		wg.Add(1)
		go func(idWord *uint, onClose func()) {

			defer onClose()

			path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				rs := scanner.Text()
				setDocuments.Add(rs)
			}

		}(idWord, func() { wg.Done() })

	}

	wg.Wait()

	for k, _ := range setDocuments.GetSet() {
		result = append(result, k)
	}

	return result
}
