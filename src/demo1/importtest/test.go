package importtest

import "fmt"

func init() {
	fmt.Println("call init")
	var array [5]int
	array[0] = 1

	arr := [5]int{1, 2, 3, 4, 5}
	for _, value := range arr {
		fmt.Println(value)
	}
}
func test() {

}
