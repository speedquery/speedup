package collection

import "sync"

type SetUint struct {
	someMapMutex sync.RWMutex
	set          map[uint]bool
}

func (self *SetUint) NewSet() *SetUint {
	self.someMapMutex = sync.RWMutex{}
	self.set = make(map[uint]bool)
	return self
}

func (self *SetUint) IsExistValue(key uint) bool {

	self.someMapMutex.Lock()
	_, exist := self.set[key]
	self.someMapMutex.Unlock()

	return exist
}

func (self *SetUint) GetSet() map[uint]bool {

	temp := make(map[uint]bool)

	self.someMapMutex.Lock()

	for key, vl := range self.set {
		temp[key] = vl
	}

	self.NewSet()

	self.someMapMutex.Unlock()

	return temp
}

func (self *SetUint) Clone() map[uint]bool {

	return self.set
}

func (self *SetUint) Add(value uint) {
	self.someMapMutex.Lock()
	self.set[value] = true
	self.someMapMutex.Unlock()
}

func (self *SetUint) Delete(value uint) {
	self.someMapMutex.Lock()
	delete(self.set, value)
	self.someMapMutex.Unlock()

}

func (st *SetUint) Size() int {
	return len(st.set)
}
