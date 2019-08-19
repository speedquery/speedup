package filesystem

import (
	"encoding/json"
	"fmt"
	"sync"
)

type AttributeWord struct {
	attributeWord map[*uint][]*uint
	someMapMutex  sync.RWMutex
}

func (attw *AttributeWord) InitAttributeWord() *AttributeWord {
	attw.someMapMutex = sync.RWMutex{}
	attw.attributeWord = make(map[*uint][]*uint)
	return attw
}

func (attw *AttributeWord) GetWordsOfAttribute(idAttribute *uint) []*uint {

	idwords := attw.attributeWord[idAttribute]
	return idwords

}

func (attw *AttributeWord) AddWordsOfAttribute(idAttribute, idWord *uint) []*uint {

	attw.someMapMutex.Lock()

	idwords, exist := attw.attributeWord[idAttribute]

	if !exist {
		idwords = make([]*uint, 0)
		idwords = append(idwords, idWord)
		attw.attributeWord[idAttribute] = idwords
	}

	existSlice := false
	for _, localIDword := range idwords {
		if localIDword == idWord {
			existSlice = true
			break
		}
	}

	if !existSlice {
		idwords = append(idwords, idWord)
		attw.attributeWord[idAttribute] = idwords
	}

	attw.someMapMutex.Unlock()

	return idwords

}

func (attw *AttributeWord) ToJson() string {

	temp := make(map[uint][]*uint)

	for key, value := range attw.attributeWord {
		temp[*key] = value
	}

	data, err := json.Marshal(temp)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
