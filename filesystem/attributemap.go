package filesystem

import (
	"encoding/json"
	"fmt"
)

type AttributeMap struct {
	attributeMap map[string]uint
	id           uint
}

//NewWordMap create new wordmap
func (att *AttributeMap) IniAttributeMap() *AttributeMap {
	att.attributeMap = make(map[string]uint)
	att.id = 0
	return att
}

//AddWord Add new word in map
func (att *AttributeMap) AddAttribute(attribute string) uint {

	value, exist := att.attributeMap[attribute]

	if !exist {
		att.id++
		value = att.id
		att.attributeMap[attribute] = value
	}

	return value
}

func (att *AttributeMap) ToJson() string {

	data, err := json.Marshal(att.attributeMap)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
