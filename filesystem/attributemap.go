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

func (self *AttributeMap) SetNewMap(newMap map[string]*uint) *AttributeMap {
	self.attributeMap = newMap
	return self
}

func (self *AttributeMap) GetPointer(id uint) *uint {
	var point *uint = nil
	for _, value := range self.attributeMap {
		if *value == id {
			point = value
		}
	}

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

	data, err := json.Marshal(self.attributeMap)

	if err != nil {
		panic(err)
	}

	return string(data)
}
