package games

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateGeometricProgression() (string, string) {
    rand.Seed(time.Now().UnixNano())
    length := rand.Intn(6) + 5 // length between 5 and 10
    start := rand.Intn(10) + 1
    ratio := rand.Intn(4) + 2

    progression := make([]int, length)
    for i := 0; i < length; i++ {
        progression[i] = start * pow(ratio, i)
    }

    hiddenIndex := rand.Intn(length)
    correctAnswer := fmt.Sprintf("%d", progression[hiddenIndex])
    progression[hiddenIndex] = -1

    var questionParts []string
    for _, value := range progression {
        if value == -1 {
            questionParts = append(questionParts, "..")
        } else {
            questionParts = append(questionParts, fmt.Sprintf("%d", value))
        }
    }

    question := strings.Join(questionParts, " ")
    return question, correctAnswer
}

func pow(base, exp int) int {
    result := 1
    for i := 0; i < exp; i++ {
        result *= base
    }
    return result
}
