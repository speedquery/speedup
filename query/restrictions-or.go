package query

type OR struct {
	orlist []Operators
}

func (self *OR) AddOR(eq Operators) *OR {

	if self.orlist == nil {
		self.orlist = make([]Operators, 0)
	}

	self.orlist = append(self.orlist, eq)

	return self

}

func (self *OR) GetListEQ() []Operators {
	return self.orlist
}
