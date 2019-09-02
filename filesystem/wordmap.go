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

func (wd *WordMap) SetNewMap(id uint, newMap map[string]*uint) *WordMap {
	wd.wordMap = newMap
	wd.id = id
	return wd
}

func (self *WordMap) GetID() uint {
	return self.id
}

//AddWord Add new word in map
func (self *WordMap) AddWord(word string) *uint {

	self.someMapMutex.Lock()

	value, exist := self.wordMap[word]

	if !exist {
		self.id++
		newvalue := self.id
		value = &newvalue
		self.wordMap[word] = value
	}

	self.someMapMutex.Unlock()

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

func (self *WordMap) ToJsonID() string {

	self.someMapMutex.Lock()

	maxID := make(map[string]uint)

	maxID["maxid"] = self.GetID()

	data, err := json.Marshal(maxID)

	if err != nil {
		panic(err)
	}

	self.someMapMutex.Unlock()

	return string(data)

}
