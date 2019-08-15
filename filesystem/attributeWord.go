package filesystem

import (
	"encoding/json"
	"fmt"
)

type AttributeWord struct {
	attributeWord map[uint][]uint
}

func (attw *AttributeWord) InitAttributeWord() *AttributeWord {
	attw.attributeWord = make(map[uint][]uint)
	return attw
}

func (attw *AttributeWord) GetWordsOfAttribute(idAttribute uint) []uint {
	idwords := attw.attributeWord[idAttribute]
	return idwords
}

func (attw *AttributeWord) AddWordsOfAttribute(idAttribute, idWord uint) []uint {

	idwords, exist := attw.attributeWord[idAttribute]

	if !exist {
		idwords = make([]uint, 0)
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

	return idwords

}

func (attw *AttributeWord) ToJson() string {

	data, err := json.Marshal(attw.attributeWord)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
