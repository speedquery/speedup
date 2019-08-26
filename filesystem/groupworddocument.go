package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
	"speedup/collection"
	"sync"
)

type GroupWordDocument struct {
	groupWordDocument map[*uint]*collection.Set
	someMapMutex      sync.RWMutex
}

func (gw *GroupWordDocument) InitGroupWordDocument() *GroupWordDocument {

	gw.someMapMutex = sync.RWMutex{}
	gw.groupWordDocument = make(map[*uint]*collection.Set)
	return gw
}

func (gw *GroupWordDocument) GetIdGroupWord(idDocument *uint) *collection.Set {

	idGroups := gw.groupWordDocument[idDocument]
	return idGroups

}

func (gw *GroupWordDocument) AddGroupWordDocument(idGroup *uint, idDocument *uint) {
	/**
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)
	**/
	if _, err := os.Stat("C:/testeGolang"); os.IsNotExist(err) {

		println("Pasta nao existe")
		os.Mkdir("C:/testeGolang", 0777)

		// path/to/whatever does not exist
	}

	/**
	gw.someMapMutex.Lock()

	idDocuments, exist := gw.groupWordDocument[idGroup]

	if !exist || idDocuments == nil {
		idDocuments = new(collection.Set).NewSet()
		idDocuments.Add(idDocument)
		gw.groupWordDocument[idGroup] = idDocuments
	} else {
		idDocuments.Add(idDocument)
	}

	gw.someMapMutex.Unlock()

	return idDocuments

	/**
		idDocuments, exist := gw.groupWordDocument[idGroup]

		if !exist || idDocuments == nil {
			idDocuments = new(collection.Set).NewSet()
			idDocuments.Add(idDocument)
			gw.groupWordDocument[idGroup] = idDocuments
		} else {
			idDocuments.Add(idDocument)
		}

		gw.someMapMutex.Unlock()

		return idDocuments
	**/
}

func (gw *GroupWordDocument) ToJson() string {

	temp := make(map[uint][]*uint)

	for key, values := range gw.groupWordDocument {

		words := make([]*uint, 0)

		for key, _ := range values.GetSet() {
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
