package query

import (
	"runtime"
	"speedup/filesystem"
	"sync"
)

func GetBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

type Query struct {
	filesystem *filesystem.FileSystem
	andList    []Operators
	orList     []*OR
}

func (self *Query) CreateQuery(filesystem *filesystem.FileSystem) *Query {
	self.filesystem = filesystem
	return self
}

func (self *Query) FilterOr(query *OR) []string {

	listEq := query.GetListEQ()

	var result []string
	for _, value := range listEq {

		result = self.FilterAnd(value)

		if len(result) > 0 {
			break
		}

	}

	return result

}

func (self *Query) Add(query Operators) *Query {

	if self.andList == nil {
		self.andList = make([]Operators, 0)
	}

	self.andList = append(self.andList, query)
	return self
}

func (self *Query) AddOR(query *OR) *Query {

	if self.andList == nil {
		self.orList = make([]*OR, 0)
	}

	self.orList = append(self.orList, query)
	return self
}

func (self *Query) GetList() []string {

	var wg sync.WaitGroup
	list := make([][]string, 0)

	for _, query := range self.andList {

		for _, eq := range query.GetList() {

			key := eq.Key
			value := eq.Value

			wg.Add(1)
			go func(key, value string, onClose func()) {

				defer onClose()

				switch query.(type) {
				case *EQ:

					result := self.FindIndexEQ(value)

					if len(result) > 0 {
						list = append(list, result)
					}

				case *NotEQ:

					result := self.FindIndexNotEQ(value)

					if len(result) > 0 {
						list = append(list, result)
					}

				}

			}(key, value, func() { wg.Done() })
		}

	}

	wg.Wait()

	if len(list) > 0 {
		result := list[0]
		for i := 1; i <= len(list)-1; i++ {
			result = difference(result, list[i])
		}

		return result
	} else {

		for _, query := range self.andList {
		}

		self.FilterOr()

	}

	return nil

}

func (self *Query) FilterAnd(query Operators) []string {

	var wg sync.WaitGroup

	list := make([][]string, 0)

	for _, eq := range query.GetList() {

		key := eq.Key
		value := eq.Value

		wg.Add(1)
		go func(key, value string, onClose func()) {

			defer onClose()

			switch query.(type) {
			case *EQ:

				result := self.FindIndexEQ(value)

				if len(result) > 0 {
					list = append(list, result)
				}

			case *NotEQ:

				result := self.FindIndexNotEQ(value)

				if len(result) > 0 {
					list = append(list, result)
				}

			}

		}(key, value, func() { wg.Done() })
	}

	wg.Wait()

	if len(list) > 0 {
		result := list[0]
		for i := 1; i <= len(list)-1; i++ {

			result = difference(result, list[i])

		}

		return result
	} else {
		return nil
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
