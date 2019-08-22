package filesystem

import (
	"container/list"
	"encoding/json"
	"fmt"
)

type AttributeGroupWord struct {
	attributeGroupWord map[*uint]*list.List
}

func (attg *AttributeGroupWord) InitAttributeGroupWord() *AttributeGroupWord {
	attg.attributeGroupWord = make(map[*uint]*list.List)
	return attg
}

/**
func (attg *AttributeGroupWord) GetIdGroupOfAttribute(idAttribute *uint) []*uint {

	idwords := attg.attributeGroupWord[idAttribute]
	return idwords

}
**/
func (attg *AttributeGroupWord) AddGroupWordsOfAttribute(idAttribute, idGroup *uint) *list.List {

	idGroups, exist := attg.attributeGroupWord[idAttribute]

	if !exist {
		idGroups = list.New()
		idGroups.PushBack(idGroup)
		attg.attributeGroupWord[idAttribute] = idGroups
	} else {
		existSlice := false

		for e := idGroups.Front(); e != nil; e = e.Next() {

			localID := e.Value.(*uint)

			if localID == idGroup {
				existSlice = true
				break
			}
		}

		if !existSlice {
			idGroups.PushBack(idGroup)
			attg.attributeGroupWord[idAttribute] = idGroups
		}
	}

	return idGroups

}

func (attg *AttributeGroupWord) ToJson() string {

	//data, err := json.Marshal(attg.attributeGroupWord)

	temp := make(map[uint][]*uint)

	for key, value := range attg.attributeGroupWord {

		groups := make([]*uint, 0)

		for e := value.Front(); e != nil; e = e.Next() {
			groups = append(groups, e.Value.(*uint))
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
