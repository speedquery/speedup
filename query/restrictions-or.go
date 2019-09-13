package query

type OR struct {
	orlist []*EQ
}

func (self *OR) AddOR(eq *EQ) *OR {

	if self.orlist == nil {
		self.orlist = make([]*EQ, 0)
	}

	self.orlist = append(self.orlist, eq)

	return self

}

func (self *OR) GetList() []*EQ {
	return self.orlist
}
