// Package entities contains support for types of
// people in the system.
package entities

// User defines a user in the program.
type User struct {
	Name  string
	email string // 小写开头 私有变量
}
