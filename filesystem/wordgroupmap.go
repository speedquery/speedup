package filesystem

import (
	"encoding/json"
	"sync"
)

type WordGroupMap struct {
	wordGroupMap map[string]*uint
	id           uint
	someMapMutex sync.RWMutex
}

//NewWordMap create new wordmap
func (wd *WordGroupMap) IniWordGroupMap() *WordGroupMap {
	wd.someMapMutex = sync.RWMutex{}
	wd.wordGroupMap = make(map[string]*uint)
	wd.id = 0
	return wd
	//att.attributeMap = make(map[string]uint)
	//att.id = 0
	//return att
}

//AddWord Add new word in map
func (wd *WordGroupMap) AddAWordGroup(wordgroup string) *uint {

	wd.someMapMutex.Lock()

	value, exist := wd.wordGroupMap[wordgroup]

	if !exist {
		wd.id++
		newvalue := wd.id
		value = &newvalue
		wd.wordGroupMap[wordgroup] = value
	}

	wd.someMapMutex.Unlock()

	return value

}

func (wd *WordGroupMap) ToJson() string {

	data, err := json.Marshal(wd.wordGroupMap)

	if err != nil {
		panic(err)
	}

	return string(data)
}
