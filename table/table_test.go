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

var data []Data = []Data{
	{BB: "bbb", AA: "aaa", dd: "DDD"},
	{"111", "a222", "333", "444"}}

func TestPrint(t *testing.T) {
	Print(data)
}

func TestPrintf(t *testing.T) {
	instead := make(map[string]string)
	instead["BB"] = "hello"
	ignore := []string{"CC"}

	fmt.Println("======= common")
	Printf(data, instead, ignore)
	fmt.Println("======= nil ignore")
	Printf(data, instead, nil)
	fmt.Println("======= nil instead")
	Printf(data, nil, ignore)
	fmt.Println("======= nil instead and ignore")
	Printf(data, nil, nil)
}
