package table

import (
	"log"
	"reflect"
)

type Operation interface {
	Do(tab *Table) Table
}

type Equal struct {
	Colume string
	Value  any
}

func (eq Equal) Do(tab *Table) Table {
	index := make([]int, 0)
	col := tab.columes[eq.Colume]
	if col == nil {
		log.Printf("The table doesn't contain colume named '%s'", eq.Colume)
		return Table{}
	}
	vType := reflect.TypeOf(eq.Value)
	if vType != col.Type {
		log.Printf("The value's type '%s' doesn't match colume's type '%s'", vType.Name(), col.Type.Name())
	}
	for i, value := range col.Data {
		if eq.Value == value.Interface() {
			index = append(index, i)
		}
	}
	columes := make(map[string]*colume)
	for i := range index {
		for name, col := range tab.columes {
			newCol := columes[name]
			if newCol == nil {
				newCol = &colume{Name: col.Name, Type: col.Type, Data: make([]reflect.Value, 0)}
				columes[name] = newCol
			}
			newCol.Data = append(newCol.Data, col.Data[i])
		}
	}
	return Table{len(index), columes}
}
