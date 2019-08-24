package filesystem

import (
	"speedup/collection"
	"sync"
)

type AttributeWord struct {
	attributeWord map[*uint]*collection.Set
	someMapMutex  sync.RWMutex
}

func (attw *AttributeWord) InitAttributeWord() *AttributeWord {
	attw.someMapMutex = sync.RWMutex{}
	attw.attributeWord = make(map[*uint]*collection.Set)
	return attw
}

/**
func (attw *AttributeWord) GetWordsOfAttribute(idAttribute *uint) []*uint {

	idwords := attw.attributeWord[idAttribute]
	return idwords

}
**/

func (attw *AttributeWord) AddWordsOfAttribute(idAttribute, idWord *uint) *collection.Set {

	attw.someMapMutex.Lock()

	idwords, exist := attw.attributeWord[idAttribute]

	if !exist || idwords == nil {
		idwords = new(collection.Set).NewSet()
		idwords.Add(idWord)
		attw.attributeWord[idAttribute] = idwords
	} else {
		idwords.Add(idWord)
	}

	attw.someMapMutex.Unlock()

	return idwords

}

/**
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
**/
