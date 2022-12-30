package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExampleData1 struct {
	AA string
	BB int
	CC bool
	dd string
	EE float32
}

type ExampleData2 struct {
	AA string
	BB int
}

var data []ExampleData1 = []ExampleData1{
	{AA: "aaa", BB: 111, CC: false, dd: "ddd", EE: 12.34},
	{"111", 222, true, "444", 1.23},
}

func TestBuildTable(t *testing.T) {
	tab := NewTable(data)
	assert.Equal(t, 4, len(tab.columes))
	assert.Equal(t, 2, tab.size)
	assert.NotNil(t, tab.columes["AA"])
	assert.NotNil(t, tab.columes["BB"])
	assert.NotNil(t, tab.columes["CC"])
	assert.NotNil(t, tab.columes["EE"])
	assert.Nil(t, tab.columes["dd"])
	tab.ShowSchema()
	tab.Show()
}

func TestSelect(t *testing.T) {
	tab := NewTable(data).Select("AA", "BB")
	assert.Equal(t, 2, len(tab.columes))
	assert.Equal(t, 2, tab.size)
	assert.NotNil(t, tab.columes["AA"])
	assert.NotNil(t, tab.columes["BB"])
	assert.Nil(t, tab.columes["CC"])
	assert.Nil(t, tab.columes["EE"])
	assert.Nil(t, tab.columes["dd"])
}

func TestUnmarshal(t *testing.T) {
	tab := NewTable(data).Select("AA", "BB")
	var results []ExampleData2
	tab.Unmarshal(&results)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, data[0].AA, results[0].AA)
	assert.Equal(t, data[0].BB, results[0].BB)
	assert.Equal(t, data[1].AA, results[1].AA)
	assert.Equal(t, data[1].BB, results[1].BB)
}

func TestEqual(t *testing.T) {
	tab := NewTable(data).Where(Equal{Condition{Colume: "BB", Value: 111}})
	tab.Show()
	var results []ExampleData1
	tab.Unmarshal(&results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, data[0].AA, results[0].AA)
	assert.Equal(t, data[0].BB, results[0].BB)
	assert.Equal(t, data[0].CC, results[0].CC)
	assert.Equal(t, data[0].EE, results[0].EE)
}
