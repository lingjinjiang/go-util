package table

type condition interface {
	evaluate(actual any) bool
	getColume() string
	getValue() any
}

type Condition struct {
	Colume string
	Value  any
}

func (c Condition) evaluate(actual any) bool {
	return false
}

func (c Condition) getColume() string {
	return c.Colume
}

func (c Condition) getValue() any {
	return c.Value
}

type Equal struct {
	Condition
}

func (eq Equal) evaluate(actual any) bool {
	return actual == eq.Value
}
