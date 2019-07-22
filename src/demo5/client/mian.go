package main

import "fmt"

func main() {

	// testMap()
	mapWithFunc()
}

func testSlice() {
	year := []string{"1", "2", "3"}
	fmt.Println(year, len(year), cap(year))
	s := make([]string, 3, 5)
	fmt.Println(s, len(s), cap(s))
	s = append(s, "1")
	fmt.Println(s, len(s), cap(s))

	s1 := s[1:3]
	fmt.Println(s1, len(s1), cap(s1))
}

func testMap() {
	colors := make(map[int]string)
	colors[1] = "bule"
	colors[2] = "yellow"
	colors[3] = "red"

	fmt.Println(len(colors), colors[1])

	for k, v := range colors {
		fmt.Println(k, v)
	}
}
func mapWithFunc() {
	funcs := make(map[string]func(n int) int)
	funcs["add"] = func(n int) int {
		return n + n
	}
	funcs["div"] = func(n int) int {
		return n / n
	}
	fmt.Println(funcs["add"](2))
}
