package filesystem

import (
	"encoding/json"
	"sync"
)

//WordMap struct
type WordMap struct {
	wordMap      map[string]*uint
	point        map[uint]*uint
	id           uint
	someMapMutex sync.RWMutex
}

//NewWordMap create new wordmap
func (self *WordMap) InitWordMap() *WordMap {
	self.someMapMutex = sync.RWMutex{}
	self.wordMap = make(map[string]*uint)
	self.point = make(map[uint]*uint)
	self.id = 0
	return self
}

func (self *WordMap) SetNewMap(id uint, newMap map[string]*uint) *WordMap {
	self.wordMap = newMap
	self.id = id

	self.point = make(map[uint]*uint)
	for _, v := range self.wordMap {
		self.point[*v] = v
	}

	return self
}

func (self *WordMap) GetID() uint {
	return self.id
}

func (self *WordMap) GetWord(value uint) string {

	self.someMapMutex.Lock()

	var vl string

	for k, v := range self.wordMap {
		if *v == value {
			vl = k
			break
		}
	}

	self.someMapMutex.Unlock()

	return vl
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
		self.point[*value] = value
	}

	self.someMapMutex.Unlock()

	return value
}

func (self *WordMap) ToJson() string {

	self.someMapMutex.Lock()

	data, err := json.Marshal(self.wordMap)

	if err != nil {
		panic(err)
	}

	self.someMapMutex.Unlock()

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
