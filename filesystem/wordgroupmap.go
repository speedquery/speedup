package filesystem

import (
	"container/list"
	"reflect"
	"sync"
)

type WordGroupMap struct {
	wordGroupMap map[*uint]*list.List
	id           uint
	someMapMutex sync.RWMutex
}

//NewWordMap create new wordmap
func (wd *WordGroupMap) IniWordGroupMap() *WordGroupMap {

	wd.someMapMutex = sync.RWMutex{}
	wd.wordGroupMap = make(map[*uint]*list.List)
	wd.id = 0
	return wd
	//att.attributeMap = make(map[string]uint)
	//att.id = 0
	//return att
}

//AddWord Add new word in map
func (wd *WordGroupMap) AddAWordGroup(wordgroup *list.List) *uint {

	wd.someMapMutex.Lock()

	exist := false

	var idgroup *uint

	for id, value := range wd.wordGroupMap {

		for e := value.Front(); e != nil; e = e.Next() {

			localID := e.Value.(*list.List)

			if reflect.DeepEqual(localID, wordgroup) {
				exist = true
				idgroup = id
				break
			}
		}

	}

	if !exist {
		wd.id++
		newvalue := wd.id
		idgroup = &newvalue
		strut := list.New()
		strut.PushBack(wordgroup)
		wd.wordGroupMap[idgroup] = strut
	}

	wd.someMapMutex.Unlock()

	return idgroup

	/**
	if !exist {
		att.id++
		value = att.id
		att.attributeMap[attribute] = value
	}

	return value
	**/
}

/**
func (wd *WordGroupMap) ToJson() string {

	temp := make(map[uint][][]*uint)

	for key, value := range wd.wordGroupMap {
		temp[*key] = value
	}

	data, err := json.Marshal(temp)

	if err != nil {
		panic(err)
	}

	return string(data)
}
**/
