package filesystem

import (
	"encoding/json"
	"sync"
)

//WordMap struct
type WordMap struct {
	wordMap      map[string]*uint
	id           uint
	someMapMutex sync.RWMutex
}

//NewWordMap create new wordmap
func (wd *WordMap) InitWordMap() *WordMap {
	wd.someMapMutex = sync.RWMutex{}
	wd.wordMap = make(map[string]*uint)
	wd.id = 0
	return wd
}

func (wd *WordMap) SetNewMap(newMap map[string]*uint) *WordMap {
	wd.wordMap = newMap
	return wd
}

//AddWord Add new word in map
func (wd *WordMap) AddWord(word string) *uint {

	wd.someMapMutex.Lock()

	value, exist := wd.wordMap[word]

	if !exist {
		wd.id++
		newvalue := wd.id
		value = &newvalue
		wd.wordMap[word] = value
	}

	wd.someMapMutex.Unlock()

	return value
}

func (wd *WordMap) ToJson() string {

	wd.someMapMutex.Lock()

	data, err := json.Marshal(wd.wordMap)

	if err != nil {
		panic(err)
	}

	wd.someMapMutex.Unlock()

	return string(data)

}
