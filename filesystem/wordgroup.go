package filesystem

import (
	"encoding/json"
	"reflect"
)

type WordGroupMap struct {
	wordGroupMap map[uint][][]uint
	id           uint
}

//NewWordMap create new wordmap
func (wd *WordGroupMap) IniWordGroupMap() *WordGroupMap {
	wd.wordGroupMap = make(map[uint][][]uint)
	wd.id = 0
	return wd
	//att.attributeMap = make(map[string]uint)
	//att.id = 0
	//return att
}

//AddWord Add new word in map
func (wd *WordGroupMap) AddAWordGroup(wordgroup []uint) uint {

	exist := false

	for _, value := range wd.wordGroupMap {

		for _, vl := range value {
			//println("compararou:", fmt.Sprintf("%v", wordgroup), fmt.Sprintf("%v", vl))
			//println("igual?", reflect.DeepEqual(vl, wordgroup))

			if reflect.DeepEqual(vl, wordgroup) {
				exist = true
				break
			}

		}

	}

	if !exist {
		wd.id += 1
		strut := make([][]uint, 0)
		strut = append(strut, wordgroup)
		wd.wordGroupMap[wd.id] = strut
	}

	return 0

	/**
	if !exist {
		att.id++
		value = att.id
		att.attributeMap[attribute] = value
	}

	return value
	**/
}

func (wd *WordGroupMap) ToJson() string {

	data, err := json.Marshal(wd.wordGroupMap)

	if err != nil {
		panic(err)
	}

	return string(data)

}
