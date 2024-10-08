package games

import (
	"fmt"
	"math/rand"
	"time"
)

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(a, b, c int) int {
    abGcd := gcd(a, b)
    abLcm := (a * b) / abGcd
    return (abLcm * c) / gcd(abLcm, c)
}

func GenerateLCMQuestion() (string, string) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num1 := r.Intn(100) + 1
	num2 := r.Intn(100) + 1
	num3 := r.Intn(100) + 1

    question := fmt.Sprintf("%d %d %d", num1, num2, num3)
    correctAnswer := fmt.Sprintf("%d", lcm(num1, num2, num3))
    return question, correctAnswer
}
