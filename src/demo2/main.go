package main

import "fmt"

func main() {
	testType()
}

func testMap() {
	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral":     "#ff7F50",
		"DarkGray":  "#a9a9a9",
	}
	colors["Red"] = "#da1337"

	fmt.Println(colors["Red"])

	value, exists := colors["Red"]
	if exists {
		fmt.Println(exists)
		fmt.Println(value)
	}
}

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func testType() {

	var bill user
	bill.name = "bill"
	bill.email = "1099263878@qq.com"
	bill.ext = 1
	bill.privileged = false

	lisa := user{
		name: "lisa",
	}

	bill.notify()
	fmt.Println(lisa)
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}
