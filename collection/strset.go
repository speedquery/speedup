package collection

import "sync"

type StrSet struct {
	someMapMutex sync.RWMutex
	set          map[string]bool
}

func (self *StrSet) NewSet() *StrSet {
	self.someMapMutex = sync.RWMutex{}
	self.set = make(map[string]bool)
	return self
}

func (self *StrSet) IsExistValue(key string) bool {

	self.someMapMutex.Lock()
	_, exist := self.set[key]
	self.someMapMutex.Unlock()

	return exist
}

func (self *StrSet) GetSet() map[string]bool {
	return self.set
}

func (self *StrSet) Add(value string) {
	self.someMapMutex.Lock()
	self.set[value] = true
	self.someMapMutex.Unlock()
}

func (self *StrSet) Delete(value string) {
	self.someMapMutex.Lock()
	delete(self.set, value)
	self.someMapMutex.Unlock()

}

func (self *StrSet) Size() int {
	return len(self.set)
}
