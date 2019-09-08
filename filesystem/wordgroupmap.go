package filesystem

import (
	"encoding/json"
	"sync"
)

type WordGroupMap struct {
	wordGroupMap      map[string]*uint
	wordGroupMapPoint map[uint]*uint

	id           uint
	someMapMutex sync.RWMutex
}

//NewWordMap create new wordmap
func (self *WordGroupMap) IniWordGroupMap() *WordGroupMap {
	self.someMapMutex = sync.RWMutex{}
	self.wordGroupMap = make(map[string]*uint)
	self.id = 0
	return self
}

func (self *WordGroupMap) SetNewMap(id uint, newMap map[string]*uint) *WordGroupMap {
	self.someMapMutex.Lock()
	self.wordGroupMapPoint = make(map[uint]*uint)
	self.wordGroupMap = newMap
	self.id = id

	func() {
		for _, value := range self.wordGroupMap {
			self.wordGroupMapPoint[*value] = value
		}
	}()

	self.someMapMutex.Unlock()

	return self
}

func (self *WordGroupMap) GetID() uint {
	return self.id
}

func (self *WordGroupMap) GetPointer(id uint) *uint {

	self.someMapMutex.Lock()
	value, _ := self.wordGroupMapPoint[id]
	self.someMapMutex.Unlock()

	return value

	/**
		var point *uint = nil

		self.someMapMutex.Lock()

		for _, value := range self.wordGroupMap {
			if *value == id {
				point = value
			}
		}

		self.someMapMutex.Unlock()

		return point
	**/
}

func (self *WordGroupMap) GetAWordGroup(wordgroup string) *uint {

	self.someMapMutex.Lock()

	value, _ := self.wordGroupMap[wordgroup]

	self.someMapMutex.Unlock()

	return value

}

//AddWord Add new word in map
func (self *WordGroupMap) AddAWordGroup(wordgroup string) *uint {

	self.someMapMutex.Lock()

	value, exist := self.wordGroupMap[wordgroup]

	if !exist {
		self.id++
		newvalue := self.id
		value = &newvalue
		self.wordGroupMap[wordgroup] = value
	}

	self.someMapMutex.Unlock()

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
