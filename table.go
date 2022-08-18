package util

import (
	"errors"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

func PrintTable[T interface{}](data []T) error {
	if len(data) <= 0 {
		return errors.New("data can not be empty")
	}
	w := tablewriter.NewWriter(os.Stdout)
	d := data[0]
	header := reflect.TypeOf(d)
	headers := make([]string, 0)
	for i := 0; i < header.NumField(); i++ {
		field := header.Field(i)
		headers = append(headers, field.Name)
	}
	w.SetHeader(headers)

	for _, d := range data {
		row := make([]string, 0)
		value := reflect.ValueOf(d)
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			row = append(row, field.String())
		}
		w.Append(row)
	}
	w.Render()
	return nil
}
