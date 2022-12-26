package sample

import (
	"bufio"
	"fmt"
	"os"
)

func TakingUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter text : ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
