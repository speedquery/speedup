package newquery

type OR struct {
	group *GROUP
}

func (self *OR) AddGroup(group *GROUP) *OR {
	self.group = group
	return self
}

func (self *OR) GetGroup() *GROUP {
	return self.group
}

type AND struct {
	group *GROUP
}

func (self *AND) AddGroup(group *GROUP) *AND {
	self.group = group
	return self
}

func (self *AND) GetGroup() *GROUP {
	return self.group
}
