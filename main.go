package main

import (
	"fmt"
	"strings"
	"time"
)

type fih struct { //fih
	Name  string
	X     int
	Y     int
	speed int
	Art   string
}

type Bubble struct { //bubb
	X         int
	Y         int
	Direction int
	Art       string
}

func main() {
	fmt.Println("Welcome to my aquarium....Bleep...Bloop...Bloop 𓆝 ⋆｡𖦹°‧")

	fih := fih{
		Name:  "Marias",
		X:     10,
		Y:     5,
		speed: 1,
		Art:   "𓆝",
	}

	fihPoses := []string{
		"𓆝",
		"𓆟",
		"𓆞",
	}

	direction := 1
	i := 0

	maxX := 32
	maxY := 7

	bubbles := []Bubble{}

	bubbleTypes := []string{
		"∘",
		"˙",
		"○",
		"˚",
		"•",
	}

    //canvas starts from here
	for {
		if i%3 == 0 {
			bubbles = append(bubbles, Bubble{
                X:         fih.X + 2,
                Y:         fih.Y,
				Direction: 1,
				Art:       bubbleTypes[i%len(bubbleTypes)], //different bubbs from fih 
			})
		}

		for j := range bubbles {
			bubbles[j].Y--
			bubbles[j].X += bubbles[j].Direction

			if bubbles[j].X >= maxX-2 {
				bubbles[j].Direction = -1
			}

			if bubbles[j].X <= 1 {
				bubbles[j].Direction = 1
			}
		}

		fih.X = fih.X + 2*fih.speed*direction

		if fih.X >= maxX-2 {
			direction = -1
		}

		if fih.X <= 0 {
			direction = 1
		}

		fih.Art = fihPoses[i%len(fihPoses)]  //fih poses

		if i > 0 {  //cleaning the console so it's not redrawing the canvas everytime and moving fih in same aquarium; beauty
			fmt.Printf("\033[%dA", maxY+2)
		}

		fmt.Println("┌" + strings.Repeat("─", maxX) + "┐") //actual fih box (Aquarium)

		for y := 1; y <= maxY; y++ {
			row := make([]rune, maxX)

			for x := range row {
				row[x] = ' '
			}

			if y == fih.Y {
				fishX := fih.X

				if fishX >= 0 && fishX < maxX {
					for x, char := range fih.Art {
						if fishX+x < maxX {
							row[fishX+x] = char
						}
					}
				}
			}

			for _, bubble := range bubbles {
				if bubble.Y == y && bubble.X >= 0 && bubble.X < maxX {
					row[bubble.X] = []rune(bubble.Art)[0]
				}
			}

			fmt.Println("│" + string(row) + "│")
		}

		fmt.Println("└" + strings.Repeat("─", maxX) + "┘")

		i++

		time.Sleep(500 * time.Millisecond)
	}
}