package filesystem

import (
	"encoding/json"
	"fmt"
)

type AttributeGroupWord struct {
	attributeGroupWord map[*uint][]*uint
}

func (attg *AttributeGroupWord) InitAttributeGroupWord() *AttributeGroupWord {
	attg.attributeGroupWord = make(map[*uint][]*uint)
	return attg
}

func (attg *AttributeGroupWord) GetIdGroupOfAttribute(idAttribute *uint) []*uint {

	idwords := attg.attributeGroupWord[idAttribute]
	return idwords

}

func (attg *AttributeGroupWord) AddGroupWordsOfAttribute(idAttribute, idGroup *uint) []*uint {

	idGroups, exist := attg.attributeGroupWord[idAttribute]

	if !exist {
		idGroups = make([]*uint, 0)
		idGroups = append(idGroups, idGroup)
		attg.attributeGroupWord[idAttribute] = idGroups
	}

	existSlice := false
	for _, localID := range idGroups {
		if localID == idGroup {
			existSlice = true
			break
		}
	}

	if !existSlice {
		idGroups = append(idGroups, idGroup)
		attg.attributeGroupWord[idAttribute] = idGroups
	}

	return idGroups

}

func (attg *AttributeGroupWord) ToJson() string {

	data, err := json.Marshal(attg.attributeGroupWord)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
