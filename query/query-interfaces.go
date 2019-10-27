package newquery

type LogicalOperator interface {
	GetGroup() *GROUP
}

type Operator interface {
	GetMap() *Map
}
