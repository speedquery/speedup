package newquery

import (
	"speedup/filesystem"
	"sync"
)

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

		result = self.FilterInGroup(logicalOperator.GetGroup())

		if len(result) > 0 {
			break
		}

	}

	return result
}

func (self *QUERY) FilterInGroup(group *GROUP) []string {

	var wg sync.WaitGroup

	list := make([][]string, 0)

	qtdOperator := 0
	qtdExist := 0

	operators := group.GetOperators()

	for _, operator := range operators {

		wg.Add(1)
		go func(operator Operator, onClose func()) {

			defer onClose()

			key := operator.GetMap().Key
			value := operator.GetMap().Value

			switch operator.(type) {
			case *EQ:

				println("EQ")
				result := self.FindAttEQ(key, value)
				qtdOperator++

				if len(result) > 0 {
					qtdExist++
					list = append(list, result)
				}

			case *NotEQ:
				println("NOT EQ")
				result := self.FindAttNotEQ(key, value)
				qtdOperator++

				if len(result) > 0 {
					qtdExist++
					list = append(list, result)
				}

			case *GT:

				println("GT")

				result := self.FindAttGT(key, value)
				qtdOperator++

				if len(result) > 0 {
					println("GT: ", result[0])
					qtdExist++
					list = append(list, result)
				}

			case *GE:

				println("GE")

				result := self.FindAttGE(key, value)
				qtdOperator++

				if len(result) > 0 {
					qtdExist++
					list = append(list, result)
				}
			}

		}(operator, func() { wg.Done() })
	}

	wg.Wait()

	if len(list) > 0 {

		println("QTD: ", qtdExist, qtdOperator)

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
