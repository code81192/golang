package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	s, _ := input.ReadString('E')

	fmt.Println(s)

}
