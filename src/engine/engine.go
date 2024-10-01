package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MaxRounds = 3

func RunGame(game func() (string, string), name string) {
    fmt.Printf("Hello, %s!\n", strings.TrimSpace(name))
    reader := bufio.NewReader(os.Stdin)

    for i := 0; i < MaxRounds; i++ {
        question, correctAnswer := game()
        fmt.Printf("Question: %s\n", question)
        fmt.Print("Your answer: ")

        answer, _ := reader.ReadString('\n')
        answer = strings.TrimSpace(answer)

        if answer == correctAnswer {
            fmt.Println("Correct!")
        } else {
            fmt.Printf("'%s' is wrong answer ;(. Correct answer was '%s'.\n", answer, correctAnswer)
            fmt.Printf("Let's try again, %s!\n", strings.TrimSpace(name))
            return
        }
    }
    fmt.Printf("Congratulations, %s!\n", strings.TrimSpace(name))
}
