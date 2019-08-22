package filesystem

import (
	"container/list"
	"encoding/json"
	"fmt"
	"sync"
)

type GroupWordDocument struct {
	groupWordDocument map[*uint]*list.List
	someMapMutex      sync.RWMutex
}

func (gw *GroupWordDocument) InitGroupWordDocument() *GroupWordDocument {

	gw.someMapMutex = sync.RWMutex{}
	gw.groupWordDocument = make(map[*uint]*list.List)
	return gw
}

func (gw *GroupWordDocument) GetIdGroupWord(idDocument *uint) *list.List {

	idGroups := gw.groupWordDocument[idDocument]
	return idGroups

}

func (gw *GroupWordDocument) AddGroupWordDocument(idGroup *uint, idDocument uint) *list.List {

	//gw.someMapMutex.Lock()

	idDocuments, exist := gw.groupWordDocument[idGroup]

	if !exist || idDocuments == nil {
		idDocuments = list.New()
		idDocuments.PushBack(idDocument)
		gw.groupWordDocument[idGroup] = idDocuments
	} else {

		//println(*idGroup, *idDocument, exist)

		existSlice := false

		for e := idDocuments.Front(); e != nil; e = e.Next() {

			localID := e.Value.(uint)

			if localID == idDocument {
				existSlice = true
				break
			}
		}

		if !existSlice {
			idDocuments.PushBack(idDocument)
			//gw.groupWordDocument[idGroup] = idDocuments
		}

	}
	//gw.someMapMutex.Unlock()

	return idDocuments

}

func (gw *GroupWordDocument) ToJson() string {

	temp := make(map[uint][]*uint)

	for key, values := range gw.groupWordDocument {

		words := make([]*uint, 0)

		for e := values.Front(); e != nil; e = e.Next() {
			vl := e.Value.(uint)
			words = append(words, &vl)
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
