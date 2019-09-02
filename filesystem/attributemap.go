package filesystem

import (
	"encoding/json"
	"sync"
)

type AttributeMap struct {
	someMapMutex sync.RWMutex
	attributeMap map[string]*uint
	id           uint
}

//NewWordMap create new wordmap
func (self *AttributeMap) IniAttributeMap() *AttributeMap {
	self.someMapMutex = sync.RWMutex{}
	self.attributeMap = make(map[string]*uint)
	self.id = 0
	return self
}

func (self *AttributeMap) GetID() uint {
	return self.id
}

func (self *AttributeMap) SetNewMap(id uint, newMap map[string]*uint) *AttributeMap {
	self.someMapMutex.Lock()
	self.attributeMap = newMap
	self.id = id
	self.someMapMutex.Unlock()
	return self
}

func (self *AttributeMap) GetPointer(id uint) *uint {

	var point *uint = nil

	self.someMapMutex.Lock()

	for _, value := range self.attributeMap {
		if *value == id {
			point = value
		}
	}

	self.someMapMutex.Unlock()

	return point

}

//AddWord Add new word in map
func (self *AttributeMap) AddAttribute(attribute string) *uint {

	self.someMapMutex.Lock()

	value, exist := self.attributeMap[attribute]

	if !exist {
		self.id++
		newvalue := self.id
		value = &newvalue
		self.attributeMap[attribute] = value
	}

	self.someMapMutex.Unlock()

	return value
}

func (self *AttributeMap) ToJson() string {

	self.someMapMutex.Lock()

	data, err := json.Marshal(self.attributeMap)

	if err != nil {
		panic(err)
	}

	self.someMapMutex.Unlock()

	return string(data)
}

func (self *AttributeMap) ToJsonID() string {

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
