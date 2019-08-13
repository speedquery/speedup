package filesystem

//WordMap struct
type WordMap struct {
	wordMap map[string]uint
	id      uint
}

//NewWordMap create new wordmap
func (wd *WordMap) InitWordMap() *WordMap {
	wd.wordMap = make(map[string]uint)
	wd.id = 0
	return wd
}

//AddWord Add new word in map
func (wd *WordMap) AddWord(word string) uint {

	value, exist := wd.wordMap[word]

	if !exist {
		wd.id++
		value = wd.id
		wd.wordMap[word] = value
	}

	return value
}
