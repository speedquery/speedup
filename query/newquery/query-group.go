package newquery

type GROUP struct {
	operators []Operator
}

func (self *GROUP) AddOperator(operator Operator) *GROUP {

	if self.operators == nil {
		self.operators = make([]Operator, 0)
	}

	self.operators = append(self.operators, operator)

	return self
}

func (self *GROUP) GetOperators() []Operator {
	return self.operators
}
