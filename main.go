package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	color = []string{
		"\033[0m",
		"\033[31m",
		"\033[32m",
		"\033[33m",
		"\033[34m",
		"\033[35m",
		"\033[36m",
		"\033[37m"}
)

func main() {
	fmt.Println("you are gay?")
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	a = strings.Replace(a, "\n", "", -1)
	for Y, y := range a {
		fmt.Print(string(color[(Y)%len(color)]), string(y))
	}
	fmt.Println(string(color[0]))
}
