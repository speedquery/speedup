package collection

import "sync"

type Set struct {
	someMapMutex sync.RWMutex
	set          map[*uint]bool
}

func (self *Set) NewSet() *Set {
	self.someMapMutex = sync.RWMutex{}
	self.set = make(map[*uint]bool)
	return self
}

func (self *Set) IsExistValue(key *uint) bool {

	self.someMapMutex.Lock()
	_, exist := self.set[key]
	self.someMapMutex.Unlock()

	return exist
}

func (self *Set) GetSet() map[*uint]bool {
	return self.set
}

func (self *Set) Add(value *uint) {
	self.someMapMutex.Lock()
	self.set[value] = true
	self.someMapMutex.Unlock()
}

func (self *Set) Delete(value *uint) {
	self.someMapMutex.Lock()
	delete(self.set, value)
	self.someMapMutex.Unlock()

}

func (st *Set) Size() int {
	return len(st.set)
}
