package filesystem

import (
	"encoding/json"
	"fmt"
	"speedup/collection"
	"sync"
)

type AttributeWord struct {
	attributeWord map[*uint]*collection.Set
	someMapMutex  sync.RWMutex
}

func (self *AttributeWord) InitAttributeWord() *AttributeWord {
	self.someMapMutex = sync.RWMutex{}
	self.attributeWord = make(map[*uint]*collection.Set)
	return self
}

func (self *AttributeWord) SetNewMap(newMap map[*uint]*collection.Set) {
	self.attributeWord = newMap
}

func (self *AttributeWord) GetValue(idAttribute *uint) (*collection.Set, bool) {

	self.someMapMutex.Lock()
	idGroups, exist := self.attributeWord[idAttribute]
	self.someMapMutex.Unlock()

	return idGroups, exist

}

/**
func (attw *AttributeWord) GetWordsOfAttribute(idAttribute *uint) []*uint {

	idwords := attw.attributeWord[idAttribute]
	return idwords
}
**/

func (self *AttributeWord) AddWordsOfAttribute(idAttribute, idWord *uint) *collection.Set {

	self.someMapMutex.Lock()

	idwords, exist := self.attributeWord[idAttribute]

	if !exist || idwords == nil {
		idwords = new(collection.Set).NewSet()
		idwords.Add(idWord)
		self.attributeWord[idAttribute] = idwords
	} else {
		idwords.Add(idWord)
	}

	self.someMapMutex.Unlock()

	return idwords

}

func (self *AttributeWord) ToJson() string {

	temp := make(map[uint][]*uint)

	for key, idwords := range self.attributeWord {

		words := make([]*uint, 0)

		for key, _ := range idwords.GetSet() {
			words = append(words, key)
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
