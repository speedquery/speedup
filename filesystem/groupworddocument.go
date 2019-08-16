package filesystem

import (
	"encoding/json"
	"fmt"
)

type GroupWordDocument struct {
	groupWordDocument map[*uint][]*uint
}

func (gw *GroupWordDocument) InitGroupWordDocument() *GroupWordDocument {
	gw.groupWordDocument = make(map[*uint][]*uint)
	return gw
}

func (gw *GroupWordDocument) GetIdGroupWord(idDocument *uint) []*uint {

	idGroups := gw.groupWordDocument[idDocument]
	return idGroups

}

func (gw *GroupWordDocument) AddGroupWordDocument(idDocument, idGroup *uint) []*uint {

	idGroups, exist := gw.groupWordDocument[idDocument]

	if !exist {
		idGroups = make([]*uint, 0)
		idGroups = append(idGroups, idGroup)
		gw.groupWordDocument[idDocument] = idGroups
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
		gw.groupWordDocument[idDocument] = idGroups
	}

	return idGroups

}

func (gw *GroupWordDocument) ToJson() string {

	tempMap := make(map[uint][]*uint)

	for key, value := range gw.groupWordDocument {
		tempMap[*key] = value
	}

	data, err := json.Marshal(tempMap)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
