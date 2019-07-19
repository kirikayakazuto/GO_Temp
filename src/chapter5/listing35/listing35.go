// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main is the entry point for the application.
func main() {
	var b bytes.Buffer

	// Write a string to the buffer.
	b.Write([]byte("Hello"))

	// Use Fprintf to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!") // 将world!打印在&b里面

	// Write the content of the Buffer to stdout.
	// 输出
	io.Copy(os.Stdout, &b)
}
