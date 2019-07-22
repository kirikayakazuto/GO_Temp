package main

import (
	"fmt"
)

// 小结, 在有缓存的chan中, 只有向chan输入大于缓存个数的元素, 才会触发chan的输出
// 例如以下函数 只有buffered中输入2个以上 元素, 才会触发str, ok := buffered 使数据
// 输出到str中, 且一旦触发, 就会将buffered中的所以元素输出(前提是已经使用close关闭了bufferde)
func main() {
	buffered := make(chan string, 2)
	go func() {
		for {
			str, ok := <-buffered
			if !ok {
				fmt.Println(`no buffer`)
				return
			}
			fmt.Println(str)
		}
	}()

	buffered <- "1"
	buffered <- "2"
	// buffered <- "3"
	// buffered <- "4"
	close(buffered)
}
