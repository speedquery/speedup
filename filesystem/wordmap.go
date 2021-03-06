package filesystem

import (
	"encoding/json"
	"sync"
)

//WordMap struct
type WordMap struct {
	id              uint
	wordMap         map[string]*uint
	invertedwordMap map[*uint]*string
	pointWordMap    map[uint]*uint
	someMapMutex    sync.RWMutex
}

//NewWordMap create new wordmap
func (self *WordMap) InitWordMap() *WordMap {
	self.someMapMutex = sync.RWMutex{}
	self.wordMap = make(map[string]*uint)
	self.pointWordMap = make(map[uint]*uint)

	self.invertedwordMap = make(map[*uint]*string)
	self.id = 0
	return self
}

func (self *WordMap) SetNewMap(id uint, newMap map[string]*uint) *WordMap {

	self.wordMap = newMap
	self.id = id

	self.invertedwordMap = make(map[*uint]*string)
	self.pointWordMap = make(map[uint]*uint)

	for k, v := range self.wordMap {

		temp := k
		self.invertedwordMap[v] = &temp
		self.pointWordMap[*v] = v

	}

	return self
}

func (self *WordMap) GetValue(key *uint) *string {

	valeu, _ := self.invertedwordMap[key]

	return valeu

}

func (self *WordMap) GetPointKey(key uint) *uint {

	self.someMapMutex.Lock()
	value, _ := self.pointWordMap[key]
	self.someMapMutex.Unlock()

	return value

}

func (self *WordMap) GetPointValue(key string) *uint {

	self.someMapMutex.Lock()
	value, _ := self.wordMap[key]
	self.someMapMutex.Unlock()

	return value

}

func (self *WordMap) GetID() uint {
	return self.id
}

/**
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
**/

//AddWord Add new word in map
func (self *WordMap) AddWord(word string) *uint {

	self.someMapMutex.Lock()

	value, exist := self.wordMap[word]

	if !exist {
		self.id++
		newvalue := self.id
		value = &newvalue

		self.wordMap[word] = value
		self.invertedwordMap[value] = &word
		self.pointWordMap[*value] = value

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
