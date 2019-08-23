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
func (att *AttributeMap) IniAttributeMap() *AttributeMap {
	att.someMapMutex = sync.RWMutex{}

	att.attributeMap = make(map[string]*uint)
	att.id = 0
	return att
}

//AddWord Add new word in map
func (att *AttributeMap) AddAttribute(attribute string) *uint {

	att.someMapMutex.Lock()

	value, exist := att.attributeMap[attribute]

	if !exist {
		att.id++
		newvalue := att.id
		value = &newvalue
		att.attributeMap[attribute] = value
	}

	att.someMapMutex.Unlock()

	return value
}

func (att *AttributeMap) ToJson() string {

	data, err := json.Marshal(att.attributeMap)

	if err != nil {
		panic(err)
	}

	return string(data)

}
