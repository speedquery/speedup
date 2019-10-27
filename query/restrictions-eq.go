package query

type EQ struct {
	eqlist []*Map
}

func (self *EQ) AddEQ(eq *Map) *EQ {

	if self.eqlist == nil {
		self.eqlist = make([]*Map, 0)
	}

	self.eqlist = append(self.eqlist, eq)

	return self

}

func (self *EQ) GetList() []*Map {
	return self.eqlist
}
