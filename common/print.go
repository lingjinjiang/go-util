package common

import (
	"fmt"
	"strings"
)

const segment string = "|"
const corner string = "+"
const split string = "-"

func BuildTableStr(headers []string, data [][]string) string {
	columeNum := len(headers)
	rowNum := len(data)
	data_format := make([]string, columeNum)
	header_format := make([]string, columeNum)
	splitLine := corner
	for i := 0; i < columeNum; i++ {
		max := len(headers[i])
		for j := 0; j < rowNum; j++ {
			if len(data[j][i]) > max {
				max = len(data[j][i])
			}
		}
		splitLine += strings.Repeat(split, max+2) + corner
		header_format[i] = fmt.Sprintf(" %%-%ds ", max)
		data_format[i] = fmt.Sprintf(" %%-%ds ", max)
	}

	tabStr := splitLine + "\n"
	tabStr += buildRowStr(header_format, headers, true) + "\n"
	tabStr += splitLine + "\n"
	for _, row := range data {
		tabStr += buildRowStr(data_format, row, false) + "\n"
	}
	tabStr += splitLine

	return tabStr
}

func buildRowStr(format []string, data []string, upper bool) string {
	rowStr := segment
	for i, f := range format {
		d := data[i]
		if upper {
			d = strings.ToUpper(d)
		}
		rowStr += fmt.Sprintf(f, d) + segment
	}
	return rowStr
}
