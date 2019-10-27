package newquery

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"speedup/collection"
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

func (self *QUERY) FindAttEQ(key, value string) []string {

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

func (self *QUERY) FindAttNotEQ(key, value string) []string {

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

func (self *QUERY) FindAttGTE(key, value string) []string {

	return self.FindAttGtGe(key, value, "GTE")

}

func (self *QUERY) FindAttGT(key, value string) []string {

	return self.FindAttGtGe(key, value, "GT")

}

func (self *QUERY) FindAttGE(key, value string) []string {

	return self.FindAttGtGe(key, value, "GE")
}

func (self *QUERY) FindAttGEE(key, value string) []string {

	return self.FindAttGtGe(key, value, "GEE")
}

func (self *QUERY) FindAttGtGe(key, value, stype string) []string {

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

		switch stype {

		case "GT":
			if wordValue <= referenceValue {
				continue
			}
		case "GTE":
			if wordValue < referenceValue {
				continue
			}
		case "GE":
			if wordValue >= referenceValue {
				continue
			}
		case "GEE":
			if wordValue > referenceValue {
				continue
			}
		}

		count++
		if count > 200 {
			count = 0
			wg.Wait()
		}

		path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "invertedworddoc" + GetBar() + fmt.Sprintf("%v", *idWord) + ".txt"

		wg.Add(1)
		go func(path string, setDocuments *collection.StrSet, onClose func()) {

			defer onClose()

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

		}(path, setDocuments, func() { wg.Done() })

	}

	wg.Wait()

	for k, _ := range setDocuments.GetSet() {

		// println("Documentos: ", k)

		result = append(result, k)
	}

	return result
}
