package games

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateGeometricProgression() (string, string) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    length := r.Intn(6) + 5 
    start := r.Intn(10) + 1
    ratio := r.Intn(4) + 2

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
