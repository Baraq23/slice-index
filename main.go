package main

import "fmt"

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	st := nbrs[0]
	if len(nbrs) > 1 {
		en := nbrs[1]
		if st > en {
			return nil
		}
		if en < 0 {
			en = getIndx(en, a)
		}
		if st < 0 {
			st = getIndx(st, a)
		}

		return a[st:en]
	}

	if st < 0 {
		st = getIndx(st, a)
		return a[st:]
	}
	return a[st:]
}

func getIndx(i int, s []string) int {
	st := len(s) - 1 + i
	return st
}
