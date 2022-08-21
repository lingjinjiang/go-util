A golang util collection for myself

* table
	- [Print](#print)
	- [Printf](#printf)

## Print

Print data in table format, it will defaultly ignore the private field.

```golang
func Print[T interface{}](data []T) error
```

Example
```golang
import "github.com/lingjinjiang/goutil/table"

type Data struct {
	AA string
	BB string
	CC string
	dd string
}

func main() {
	var data []Data = []Data{
		{AA: "aaa", BB: "bbb", dd: "ddd"},
		{"111", "222", "333", "444"},
	}
	table.Print(data)
}

```
The result will be like below
```
+-----+-----+-----+
| AA  | BB  | CC  |
+-----+-----+-----+
| aaa | bbb |     |
| 111 | 222 | 333 |
+-----+-----+-----+
```

## Printf

Print data in table format. You can use custom headers instead of defaults, or ignore given columns

```golang
func Printf[T any](objs []T, instead map[string]string, ignores []string) error
```

Example
```golang
import (
	"fmt"
	"github.com/lingjinjiang/goutil/table"
)

type Data struct {
	AA string
	BB string
	CC string
	dd string
}

func main() {
	var data []Data = []Data{
		{AA: "aaa", BB: "bbb", dd: "ddd"},
		{"111", "222", "333", "444"},
	}

	instead := make(map[string]string)
	instead["BB"] = "hello"

	ignore := []string{"CC"}

	fmt.Println("======= with instead and ignore")
	table.Printf(data, instead, ignore)
	fmt.Println("======= with instead")
	table.Printf(data, instead, nil)
	fmt.Println("======= with ignore")
	table.Printf(data, nil, ignore)
	fmt.Println("======= without instead or ignore")
	table.Printf(data, nil, nil)
}

```

The result will be like below
```
======= with instead and ignore
+-----+-------+-----+
| AA  | HELLO | DD  |
+-----+-------+-----+
| aaa | bbb   | ddd |
| 111 |   222 | 444 |
+-----+-------+-----+
======= with instead
+-----+-------+-----+-----+
| AA  | HELLO | CC  | DD  |
+-----+-------+-----+-----+
| aaa | bbb   |     | ddd |
| 111 |   222 | 333 | 444 |
+-----+-------+-----+-----+
======= with ignore
+-----+-----+-----+
| AA  | BB  | DD  |
+-----+-----+-----+
| aaa | bbb | ddd |
| 111 | 222 | 444 |
+-----+-----+-----+
======= without instead or ignore
+-----+-----+-----+-----+
| AA  | BB  | CC  | DD  |
+-----+-----+-----+-----+
| aaa | bbb |     | ddd |
| 111 | 222 | 333 | 444 |
+-----+-----+-----+-----+
```
