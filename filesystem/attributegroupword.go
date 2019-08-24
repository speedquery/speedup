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

/**
func (attg *AttributeGroupWord) GetIdGroupOfAttribute(idAttribute *uint) []*uint {

	idwords := attg.attributeGroupWord[idAttribute]
	return idwords

}
**/
func (attg *AttributeGroupWord) AddGroupWordsOfAttribute(idAttribute, idGroup *uint) {

	attg.someMapMutex.Lock()

	idGroups, exist := attg.attributeGroupWord[idAttribute]

	if idGroups == nil || !exist {
		idGroups = new(collection.Set).NewSet()
		idGroups.Add(idGroup)
		attg.attributeGroupWord[idAttribute] = idGroups
	} else {
		idGroups.Add(idGroup)
	}

	attg.someMapMutex.Unlock()

	//return idGroups

}

func (attg *AttributeGroupWord) ToJson() string {

	//data, err := json.Marshal(attg.attributeGroupWord)

	temp := make(map[uint][]*uint)

	for key, value := range attg.attributeGroupWord {

		groups := make([]*uint, 0)

		for key, _ := range value.GetSet() {
			groups = append(groups, key)
		}

		temp[*key] = groups
	}

	data, err := json.Marshal(temp)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
