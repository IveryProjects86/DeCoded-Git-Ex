// 514 class
package main

import (
"fmt"
"bufio"
"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("What is your age?: ")
	age, err := buf.ReadBytes('\n')

	if err != nil {
		fmt.Println(age)
	} else {
		fmt.Println(err)
	}
}
// Nil is coming up, something small is wrong but you get the concept. Input.go needed more than input.py. Don't need to know tthat just practicing concept 