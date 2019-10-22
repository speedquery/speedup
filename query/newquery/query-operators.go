package newquery

type Map struct {
	Key   string
	Value string
}

type EQ struct {
	opmap *Map
}

func (self *EQ) Add(opmap *Map) *EQ {
	self.opmap = opmap
	return self
}

func (self *EQ) GetMap() *Map {
	return self.opmap
}

func (self *EQ) GetTypeName() string {
	return "EQ"
}

type NotEQ struct {
	opmap *Map
}

func (self *NotEQ) Add(opmap *Map) *NotEQ {
	self.opmap = opmap
	return self
}

func (self *NotEQ) GetMap() *Map {
	return self.opmap
}

func (self *NotEQ) GetTypeName() string {
	return "NotEQ"
}

type GT struct {
	opmap *Map
}

func (self *GT) Add(opmap *Map) *GT {
	self.opmap = opmap
	return self
}

func (self *GT) GetMap() *Map {
	return self.opmap
}

func (self *GT) GetTypeName() string {
	return "GT"
}

type GE struct {
	opmap *Map
}

func (self *GE) Add(opmap *Map) *GE {
	self.opmap = opmap
	return self
}

func (self *GE) GetMap() *Map {
	return self.opmap
}

func (self *GE) GetTypeName() string {
	return "GE"
}
