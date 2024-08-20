## slice
- The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.
- The integers can be negative.

#### Example

```bash
a := []string{"coding", "algorithm", "ascii", "package", "golang"}
fmt.Printf("%#v\n", piscine.Slice(a, 1))
fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
fmt.Printf("%#v\n", piscine.Slice(a, -3))
fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))

```

_Output_

```bash
$ go run .
[]string{"algorithm", "ascii", "package", "golang"}
[]string{"ascii", "package"}
[]string{"ascii", "package", "golang"}
[]string{"package"}
[]string(nil)
```

### Notion
- slice
- slice indexing

### Author
- [Barrack Otieno](https://www.linkedin.com/in/barrack-kope-otieno-064a43244)