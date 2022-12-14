package table

import (
	"fmt"
	"testing"
)

type Data struct {
	AA string
	BB int
	CC bool
	dd string
	EE float32
}

var data []Data = []Data{
	{AA: "aaa", BB: 111, CC: false, dd: "ddd", EE: 12.34},
	{"111", 222, true, "444", 1.23},
}

func TestPrint(t *testing.T) {

	Print(data)
}

func TestPrintf(t *testing.T) {

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

func TestBuildTable(t *testing.T) {
	tab := NewTable(data)
	tab.ShowSchema()
	tab.Show()
}

func TestSelect(t *testing.T) {
	tab := NewTable(data)
	tab2 := tab.Select("AA", "dd")
	tab2.Show()
}
