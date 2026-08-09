package main

import ("fmt"
		"time"
	"strings")
type Fish struct {
	Name string
	X int
	Y int
	speed int
	Art string
}

func main(){
	fmt.Println("Welcome to my aquarium....Bleep...Bloop...Bloop 𓆝 ⋆｡𖦹°‧ ")

	fish := Fish{
    Name:  "Marias",
    X:     10,
    Y:     5,
    speed: 1,
	Art : "𓆝",
}
fishPoses := []string{
    "𓆝",
    "𓆟",
    "𓆞",
}

direction := 1
i := 0
maxX := 28

for {
    fish.X = fish.X + 2*fish.speed*direction

    if fish.X >= maxX {
        direction = -1
    }

    if fish.X <= 0 {
        direction = 1
    }

    fish.Art = fishPoses[i%len(fishPoses)]

    if i > 0 {
        fmt.Print("\033[6A")
    }

    fmt.Println("┌──────────────────────────────────┐")
    fmt.Println("│                                  │")
    fmt.Println("│" + strings.Repeat(" ", fish.X) + fish.Art + strings.Repeat(" ", maxX-fish.X))
    fmt.Println("│                                  │")
    fmt.Println("│                                  │")
    fmt.Println("└──────────────────────────────────┘")

    i++

    time.Sleep(500 * time.Millisecond)
}
}