package сli

import (
	"bufio"
	"fmt"
	"os"
)

func WelcomeUser() string {
    fmt.Println("Welcome to the Brain Games!")
    fmt.Print("May I have your name? ")

    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    return name
}