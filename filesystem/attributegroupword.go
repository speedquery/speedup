package filesystem

import (
	"encoding/json"
	"fmt"
	"speedup/collection"
	"sync"
)

type AttributeGroupWord struct {
	attributeGroupWord map[*uint]*collection.Set
	someMapMutex       sync.RWMutex
}

func (attg *AttributeGroupWord) InitAttributeGroupWord() *AttributeGroupWord {
	attg.someMapMutex = sync.RWMutex{}
	attg.attributeGroupWord = make(map[*uint]*collection.Set)
	return attg
}

func (attg *AttributeGroupWord) SetNewMap(newMap map[*uint]*collection.Set) *AttributeGroupWord {
	attg.attributeGroupWord = newMap
	return attg
}

/**
func (attg *AttributeGroupWord) GetIdGroupOfAttribute(idAttribute *uint) []*uint {

	idwords := attg.attributeGroupWord[idAttribute]
	return idwords

}
**/

func (self *AttributeGroupWord) GetValue(idAttribute *uint) (*collection.Set, bool) {

	self.someMapMutex.Lock()
	idGroups, exist := self.attributeGroupWord[idAttribute]
	self.someMapMutex.Unlock()

	return idGroups, exist

}

func (self *AttributeGroupWord) SetValue(idAttribute *uint, idGroups *collection.Set) {

	self.someMapMutex.Lock()
	self.attributeGroupWord[idAttribute] = idGroups
	self.someMapMutex.Unlock()

}

func (self *AttributeGroupWord) AddGroupWordsOfAttribute(idAttribute, idGroup *uint) {

	idGroups, exist := self.GetValue(idAttribute)

	if idGroups == nil || !exist {
		idGroups = new(collection.Set).NewSet()
		idGroups.Add(idGroup)
		self.SetValue(idAttribute, idGroups)

	} else {
		idGroups.Add(idGroup)
	}

}

func (self *AttributeGroupWord) ToJson() string {

	//data, err := json.Marshal(attg.attributeGroupWord)

	temp := make(map[uint][]*uint)

	self.someMapMutex.Lock()

	for key, value := range self.attributeGroupWord {

		groups := make([]*uint, 0)

		for key, _ := range value.GetSet() {
			groups = append(groups, key)
		}

		temp[*key] = groups
	}

	self.someMapMutex.Unlock()

	data, err := json.Marshal(temp)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
