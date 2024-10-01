package main

import (
	"fmt"
	cli "projects-mrpo/src/cli"
	"projects-mrpo/src/engine"
	"projects-mrpo/src/games"
)

func main() {
    name := cli.WelcomeUser()

    fmt.Println("Let's play LCM game!")
    engine.RunGame(games.GenerateLCMQuestion, name)

    fmt.Println("\nNow let's play Progression game!")
    engine.RunGame(games.GenerateGeometricProgression, name)
}
