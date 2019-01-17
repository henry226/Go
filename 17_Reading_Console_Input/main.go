package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell (0 to exit)")
	fmt.Println("------------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello world")
		} else if strings.Compare("0", text) == 0 {
			fmt.Println("Bye!")
			break
		}
	}
}
