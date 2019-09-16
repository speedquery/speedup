package query

type NotEQ struct {
	eqlist []*Map
}

func (self *NotEQ) AddEQ(eq *Map) *NotEQ {

	if self.eqlist == nil {
		self.eqlist = make([]*Map, 0)
	}

	self.eqlist = append(self.eqlist, eq)

	return self

}

func (self *NotEQ) GetList() []*Map {
	return self.eqlist
}
