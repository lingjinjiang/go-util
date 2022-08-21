package table

import (
	"fmt"
	"testing"
)

type Data struct {
	AA string
	BB string
	CC string
	dd string
}

func TestPrint(t *testing.T) {
	var data []Data = []Data{
		{AA: "aaa", BB: "bbb", dd: "ddd"},
		{"111", "222", "333", "444"},
	}
	Print(data)
}

func TestPrintf(t *testing.T) {
	var data []Data = []Data{
		{AA: "aaa", BB: "bbb", dd: "ddd"},
		{"111", "222", "333", "444"},
	}

	instead := make(map[string]string)
	instead["BB"] = "hello"

	ignore := []string{"CC"}

	fmt.Println("======= with instead and ignore")
	Printf(data, instead, ignore)
	fmt.Println("======= with instead")
	Printf(data, instead, nil)
	fmt.Println("======= with ignore")
	Printf(data, nil, ignore)
	fmt.Println("======= without instead or ignore")
	Printf(data, nil, nil)
}
