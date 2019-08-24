package collection

type Set struct {
	set map[*uint]bool
}

func (st *Set) NewSet() *Set {
	st.set = make(map[*uint]bool)
	return st
}

func (st *Set) Add(value *uint) {
	st.set[value] = true
}

func (st *Set) Delete(value *uint) {
	delete(st.set, value)
}

func (st *Set) Size() int {
	return len(st.set)
}
