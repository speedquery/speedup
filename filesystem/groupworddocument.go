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

func (gw *GroupWordDocument) AddGroupWordDocument(idGroup, idDocument *uint) []*uint {

	idDocuments, exist := gw.groupWordDocument[idGroup]

	if !exist {
		idDocuments = make([]*uint, 0)
		idDocuments = append(idDocuments, idDocument)
		gw.groupWordDocument[idGroup] = idDocuments
	}

	//println(*idGroup, *idDocument, exist)

	existSlice := false
	for _, localID := range idDocuments {
		if localID == idDocument {
			existSlice = true
			break
		}
	}

	if !existSlice {
		idDocuments = append(idDocuments, idDocument)
		gw.groupWordDocument[idGroup] = idDocuments
	}

	return idDocuments

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
