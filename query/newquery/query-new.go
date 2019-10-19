package newquery

import (
	"speedup/filesystem"
	"sync"
)

type Map struct {
	Key   string
	Value string
}

type Operator interface {
	GetMap() *Map
}

/**********************/

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

type LogicalOperator interface {
	GetGroup() *GROUP
}

///////////////////////////

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

/////////////////////////////

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

func (self *GROUP) GetOperators(operator Operator) []Operator {
	return self.operators
}

//////////////

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

//////////////
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

type QUERY struct {
	logicalOperator []LogicalOperator
	filesystem      *filesystem.FileSystem
}

func (self *QUERY) Create(filesystem *filesystem.FileSystem) *QUERY {
	self.filesystem = filesystem
	return self
}

func (self *QUERY) Add(logicalOperator LogicalOperator) *QUERY {

	if self.logicalOperator == nil {
		self.logicalOperator = make([]LogicalOperator, 0)
	}

	self.logicalOperator = append(self.logicalOperator, logicalOperator)

	return self

}

func (self *QUERY) GetList() []string {

	var result []string

	for _, logicalOperator := range self.logicalOperator {

		switch logicalOperator.(type) {
		case *AND:
			result = self.FilterAnd(logicalOperator.GetGroup())
		case *OR:
			println("OR")
		}
	}

	return result
}

func (self *QUERY) FilterAnd(group *GROUP) []string {

	var wg sync.WaitGroup

	list := make([][]string, 0)

	qtdOperator := 0
	qtdExist := 0

	for _, operator := range group.operators {

		key := operator.GetMap().Key
		value := operator.GetMap().Value

		wg.Add(1)
		go func(key, value string, onClose func()) {

			defer onClose()

			switch operator.(type) {
			case *EQ:

				result := self.FindAttEQ(key, value)
				qtdOperator++

				if len(result) > 0 {
					qtdExist++
					list = append(list, result)
				}

			case *NotEQ:

				result := self.FindAttNotEQ(key, value)
				qtdOperator++

				if len(result) > 0 {
					qtdExist++
					list = append(list, result)
				}

			}

		}(key, value, func() { wg.Done() })
	}

	wg.Wait()

	if len(list) > 0 {

		if qtdExist != qtdOperator {
			return make([]string, 0)
		}

		result := list[0]
		for i := 1; i <= len(list)-1; i++ {

			result = difference(result, list[i])

		}
		return result
	} else {
		return make([]string, 0)
	}

}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; found {
			diff = append(diff, x)
		}
	}
	return diff
}
