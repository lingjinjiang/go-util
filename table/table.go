package table

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

type Table struct {
	columes map[string]*colume
}

type colume struct {
	Name string
	Type reflect.Type
	Data []reflect.Value
}

func Print[T any](objs []T) error {
	if len(objs) <= 0 {
		return errors.New("data can not be empty")
	}
	o := objs[0]
	objType := reflect.TypeOf(o)
	ignores := make([]string, 0)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if !field.IsExported() {
			ignores = append(ignores, field.Name)
		}
	}

	headers, data := buildHeaderAndData(objs, nil, ignores)

	return print(headers, data)
}

func Printf[T any](objs []T, instead map[string]string, ignores []string) error {
	if len(objs) <= 0 {
		return errors.New("objs can not be empty")
	}

	headers, data := buildHeaderAndData(objs, instead, ignores)

	return print(headers, data)
}

func buildHeaderAndData[T any](objs []T, instead map[string]string, ignores []string) ([]string, [][]string) {
	o := objs[0]
	objType := reflect.TypeOf(o)
	ignoreIndex := make(map[int]bool)
	headerIndex := make(map[string]int)
	for i := 0; i < objType.NumField(); i++ {
		headerIndex[objType.Field(i).Name] = i
		ignoreIndex[i] = false
	}

	if len(ignores) > 0 {
		for _, ignore := range ignores {
			index := headerIndex[ignore]
			if index >= 0 {
				ignoreIndex[index] = true
			}
		}
	}

	headers := make([]string, 0)

	for i := 0; i < objType.NumField(); i++ {
		if ignoreIndex[i] {
			continue
		}
		header := objType.Field(i).Name
		if len(instead) > 0 && len(instead[header]) > 0 {
			headers = append(headers, instead[header])
		} else {
			headers = append(headers, header)
		}
	}

	data := make([][]string, 0)
	for _, o := range objs {
		row := make([]string, 0)
		values := reflect.ValueOf(o)
		for i := 0; i < values.NumField(); i++ {
			if ignoreIndex[i] {
				continue
			}
			row = append(row, fmt.Sprint(values.Field(i)))
		}
		data = append(data, row)
	}

	return headers, data
}

func print(headers []string, data [][]string) error {
	if len(headers) <= 0 {
		return errors.New("no available data print")
	}

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader(headers)

	for _, row := range data {
		w.Append(row)
	}
	w.Render()
	return nil
}

func NewTable[T any](objs []T) Table {
	o := objs[0]
	objType := reflect.TypeOf(o)
	headerIndex := make(map[int]string)
	columes := make(map[string]*colume)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		col := colume{
			Name: field.Name,
			Type: field.Type,
			Data: make([]reflect.Value, 0),
		}
		columes[col.Name] = &col
		headerIndex[i] = col.Name
	}

	for _, o := range objs {
		values := reflect.ValueOf(o)
		for i := 0; i < values.NumField(); i++ {
			col := columes[headerIndex[i]]
			col.Data = append(col.Data, values.Field(i))
		}
	}

	return Table{columes: columes}
}

func (tab *Table) ShowSchema() {
	data := make([][]string, 0)
	for _, col := range tab.columes {
		data = append(data, []string{col.Name, col.Type.Name()})
	}
	print([]string{"name", "type"}, data)
}

func (tab *Table) Show() {
	data := make([][]string, len(tab.columes))
	header := make([]string, 0)
	for _, col := range tab.columes {
		header = append(header, col.Name)
		for i, d := range col.Data {
			if data[i] == nil {
				data[i] = make([]string, 0)
			}
			data[i] = append(data[i], fmt.Sprint(d))
		}
	}

	print(header, data)
}
