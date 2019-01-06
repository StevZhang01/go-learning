// 利用map实现字符串集合
// dedup.go使用map的键来存储已经出现过的行
// 来确保接下来出现的相同行不会输出
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var seen = make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
