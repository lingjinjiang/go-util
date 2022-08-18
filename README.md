A golang util collection for myself

* [PrintTable](#printtable)

## PrintTable

```golang
func PrintTable[T interface{}](data []T) error
```
Print data in table format

Example
```golang
type Data struct {
	A string
	B string
	C string
}

func main() {
	d1 := Data{"1", "2", "3"}
	d2 := Data{"4", "5", "6"}
	data := []Data{d1, d2}

	PrintTable(data)
}

```
The result will be like below
```
+---+---+---+
| A | B | C |
+---+---+---+
| 1 | 2 | 3 |
| 4 | 5 | 6 |
+---+---+---+
```
