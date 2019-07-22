package main

import (
	"fmt"
	"testing"
)

func main() {

	// testMap()
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

//TestDownload validates the http Get function can download content.
func TestDownload(t *testing.T) {
	users := make(map[string]string)
	users["dl"] = "denglang"

	t.Log(len(users))
}

func TestMap(t *testing.T) {
	classes := make(map[string]func(n int) int)
	classes["add"] = func(n int) int {
		return n + n
	}
	classes["mult"] = func(n int) int {
		return n * n
	}
	t.Log(classes["add"](2), classes["mult"](3))
}

// 切片
func TestSlice2(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := s[1:]
	t.Log(s1)
	s1 = append(s1, 4, 5)
	t.Log(s1)
}
