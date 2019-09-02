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
}

func (self *WordGroupMap) SetNewMap(id uint, newMap map[string]*uint) *WordGroupMap {
	self.someMapMutex.Lock()
	self.wordGroupMap = newMap
	self.id = id
	self.someMapMutex.Unlock()

	return self
}

func (self *WordGroupMap) GetID() uint {
	return self.id
}

func (self *WordGroupMap) GetPointer(id uint) *uint {

	var point *uint = nil

	self.someMapMutex.Lock()

	for _, value := range self.wordGroupMap {
		if *value == id {
			point = value
		}
	}

	self.someMapMutex.Unlock()

	return point

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

func (self *WordGroupMap) ToJson() string {

	self.someMapMutex.Lock()

	data, err := json.Marshal(self.wordGroupMap)

	if err != nil {
		panic(err)
	}

	self.someMapMutex.Unlock()

	return string(data)
}

func (self *WordGroupMap) ToJsonID() string {

	self.someMapMutex.Lock()

	maxId := make(map[string]uint)

	maxId["maxid"] = self.GetID()

	data, err := json.Marshal(maxId)

	if err != nil {
		panic(err)
	}

	self.someMapMutex.Unlock()

	return string(data)
}
