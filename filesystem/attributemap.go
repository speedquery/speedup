package filesystem

import (
	"encoding/json"
	"sync"
)

type AttributeMap struct {
	someMapMutex      sync.RWMutex
	attributeMap      map[string]*uint
	attributeMapPoint map[uint]*uint
	id                uint
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
	self.attributeMapPoint = make(map[uint]*uint)
	self.attributeMap = newMap
	self.id = id

	func() {
		for _, value := range self.attributeMap {
			self.attributeMapPoint[*value] = value
		}
	}()

	self.someMapMutex.Unlock()
	return self
}

func (self *AttributeMap) GetPointer(id uint) *uint {
	value, _ := self.attributeMapPoint[id]
	return value
}

func (self *AttributeMap) GetAttribute(attribute string) *uint {

	self.someMapMutex.Lock()
	value, _ := self.attributeMap[attribute]
	self.someMapMutex.Unlock()

	return value

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
