package collection

import "container/list"

type SetList struct {
	set map[*list.List]bool
}

func (st *SetList) GetSetList() map[*list.List]bool {
	return st.set
}

func (st *SetList) NewSetList() *SetList {
	return st
}

func (st *SetList) Add(value *list.List) {
	st.set[value] = true
}

func (st *SetList) Delete(value *list.List) {
	delete(st.set, value)
}

func (st *SetList) Size() int {
	return len(st.set)
}
