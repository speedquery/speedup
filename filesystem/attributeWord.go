package filesystem

import (
	"container/list"
	"encoding/json"
	"fmt"
	"sync"
)

type AttributeWord struct {
	attributeWord map[*uint]*list.List
	someMapMutex  sync.RWMutex
}

func (attw *AttributeWord) InitAttributeWord() *AttributeWord {
	attw.someMapMutex = sync.RWMutex{}
	attw.attributeWord = make(map[*uint]*list.List)
	return attw
}

/**
func (attw *AttributeWord) GetWordsOfAttribute(idAttribute *uint) []*uint {

	idwords := attw.attributeWord[idAttribute]
	return idwords

}
**/

func (attw *AttributeWord) AddWordsOfAttribute(idAttribute, idWord *uint) *list.List {

	//attw.someMapMutex.Lock()
	idwords, exist := attw.attributeWord[idAttribute]

	//println("Nil?", idwords == nil)

	if !exist || idwords == nil {
		idwords = list.New()
		idwords.PushBack(idWord)
		attw.attributeWord[idAttribute] = idwords
	} else {
		existSlice := false
		for e := idwords.Front(); e != nil; e = e.Next() {

			localIDword := e.Value.(*uint)

			if localIDword == idWord {
				//println("Existe slice?", *idWord)
				existSlice = true
				break
			}
		}

		if !existSlice {
			idwords.PushBack(idWord)
			attw.attributeWord[idAttribute] = idwords
		}

	}

	return idwords

}

func (attw *AttributeWord) ToJson() string {

	temp := make(map[uint][]*uint)

	for key, idwords := range attw.attributeWord {

		words := make([]*uint, 0)

		for e := idwords.Front(); e != nil; e = e.Next() {
			words = append(words, e.Value.(*uint))
		}

		temp[*key] = words
	}

	data, err := json.Marshal(temp)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
