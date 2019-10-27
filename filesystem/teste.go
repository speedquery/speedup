package filesystem

/**
import (
	"fmt"
	"os"
	"speedup/collection"
	"sync"
)

type GroupWordDocument struct {
	groupWordDocument *collection.Set
	someMapMutex      sync.RWMutex
	folder            string
}

func (gw *GroupWordDocument) InitGroupWordDocument() *GroupWordDocument {

	gw.someMapMutex = sync.RWMutex{}
	gw.groupWordDocument = new(collection.Set).NewSet()

	gw.folder = "/users/thiagorodrigues/documents/goteste"

	if _, err := os.Stat(gw.folder); os.IsNotExist(err) {
		os.Mkdir(gw.folder, 0777)
	}

	return gw
}

func (gw *GroupWordDocument) createFile(name uint) {

	if _, err := os.Stat(gw.folder + "//" + fmt.Sprintf("%v", name) + ".txt"); os.IsNotExist(err) {

		file, err := os.Create(gw.folder + "//" + fmt.Sprintf("%v", name) + ".txt")

		if err != nil {
			panic(err)
		}

		defer file.Close()
	}

}

/**
func (gw *GroupWordDocument) GetIdGroupWord(idDocument *uint) *collection.Set {

	idGroups := gw.groupWordDocument[idDocument]
	return idGroups

}

func (gw *GroupWordDocument) AddGroupWordDocument(idGroup *uint, idDocument *uint) {

	exist := gw.groupWordDocument.IsExistValue(idGroup)

	if !exist {
		gw.createFile(*idGroup)
	}

	f, err := os.OpenFile(gw.folder+"//"+fmt.Sprintf("%v", *idGroup)+".txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	str := fmt.Sprintf("%v", *idDocument)

	if _, err = f.WriteString(str + "\n"); err != nil {
		panic(err)
	}

	/**
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)
	**/

// path/to/whatever does not exist

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
**/
