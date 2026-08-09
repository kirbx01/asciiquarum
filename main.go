package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"asciiquarum/audio"
	"asciiquarum/ui"
)

type fish struct {
	Name      string
	X         int
	Y         int
	DX        int
	DY        int
	Speed     int
	Art       string
	Bounce    bool
	MoveEvery int
	Frame     int
	Alive     bool
	Render    func(...string) string
}

type bubble struct {
	X         int
	Y         int
	Direction int
	Art       string
}


type sparkle struct {
	X         int
	Y         int
	Frames    []string
	Phase     int
	Speed     int
	Direction int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	maxX := 64
	maxY := 14

	fishes := []fish{
		{
			Name:      "Marias",
			X:         10,
			Y:         5,
			DX:        1,
			DY:        0,
			Speed:     1,
			Art:       "𓆝",
			Bounce:    false,
			MoveEvery: 1,
			Alive:     true,
			Render:    ui.Fih.Render,
		},
		{
			Name:      "Tubi",
			X:         4,
			Y:         8,
			DX:        1,
			DY:        1,
			Speed:     1,
			Art:       "𓆛",
			Bounce:    true,
			MoveEvery: 1,
			Alive:     true,
			Render:    ui.Fih.Render,
		},
		{
			Name:      "Gala",
			X:         20,
			Y:         3,
			DX:        -1,
			DY:        1,
			Speed:     1,
			Art:       "𓆡",
			Bounce:    true,
			MoveEvery: 1,
			Alive:     true,
			Render:    ui.Fih.Render,
		},
		{
			Name:      "Ari",
			X:         12,
			Y:         11,
			DX:        1,
			DY:        -1,
			Speed:     1,
			Art:       "𓆉",
			Bounce:    true,
			MoveEvery: 2,
			Alive:     true,
			Render:    ui.Fih.Render,
		},
		{
			Name:      "Star",
			X:         28,
			Y:         10,
			DX:        -1,
			DY:        0,
			Speed:     1,
			Art:       "𓇼",
			Bounce:    true,
			MoveEvery: 4,
			Alive:     true,
			Render:    ui.Sea.Render,
		},
	}
	sparkles := []sparkle{
		{X: 6, Y: 2, Frames: []string{".", "✦", "*", " "}, Speed: 2, Direction: 1},
		{X: 18, Y: 3, Frames: []string{".", "✦", "*", " "}, Speed: 3, Direction: -1},
		{X: 12, Y: 5, Frames: []string{".", "✦", "*", " "}, Speed: 4, Direction: 1},
	}

	bubbles := []bubble{}

	bubbleTypes := []string{
		"∘",
		"˙",
		"○",
		"˚",
		"•",
	}

	welcome := ui.Text.Render("Welcome to my aquarium....Bleep...Bloop...Bloop 𓆝 ⋆｡𖦹°‧")
	musicStatus := ui.Music.Render("𑣲𝄞 music playing")

	go audio.AutoPlay()

	fmt.Println(welcome)

	for i := 0; ; i++ {

		if i%3 == 0 {
			for idx := range fishes {
				if fishes[idx].Alive {
					bubbles = append(bubbles, bubble{
						X:         fishes[idx].X + 2,
						Y:         fishes[idx].Y,
						Direction: 1,
						Art:       bubbleTypes[(i+idx)%len(bubbleTypes)],
					})
				}
			}
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

		for idx := range fishes {
			fish := &fishes[idx]

			if !fish.Alive || fish.MoveEvery <= 0 {
				continue
			}

			if i%fish.MoveEvery != 0 {
				continue
			}

			fish.X += fish.DX * fish.Speed
			fish.Y += fish.DY * fish.Speed
			fish.Frame++

			if !fish.Bounce {
				if fish.X < 0 || fish.X >= maxX {
					fish.X = -len([]rune(fish.Art))
				}

				continue
			}

			if fish.X >= maxX-1 {
				fish.DX = -1
			}
			if fish.X <= 1 {
				fish.DX = 1
			}

			if fish.Y >= maxY {
				fish.DY = -1
			}
			if fish.Y <= 1 {
				fish.DY = 1
			}
		}


		for idx := range sparkles {
			if i%sparkles[idx].Speed == 0 {
				sparkles[idx].Phase =
					(sparkles[idx].Phase + 1) % len(sparkles[idx].Frames)

				sparkles[idx].X += sparkles[idx].Direction

				if sparkles[idx].X <= 1 || sparkles[idx].X >= maxX-2 {
					sparkles[idx].Direction *= -1
				}
			}
		}

		if i > 0 {
			fmt.Printf("\033[%dA", maxY+3)
		}

		fmt.Println(
			ui.Border.Render(
				"┌" + strings.Repeat("─", maxX) + "┐",
			),
		)

		for y := 1; y <= maxY; y++ {
			row := make([]string, maxX)

			for x := range row {
				row[x] = " "
			}

			for _, sparkle := range sparkles {
				if sparkle.Y == y &&
					sparkle.X >= 0 &&
					sparkle.X < maxX {

					row[sparkle.X] =
						ui.Sea.Render(sparkle.Frames[sparkle.Phase])
				}
			}

			for _, fish := range fishes {
				if !fish.Alive || fish.Y != y {
					continue
				}

				for idx, char := range []rune(fish.Art) {
					pos := fish.X + idx

					if pos >= 0 && pos < maxX {
						row[pos] = fish.Render(string(char))
					}
				}
			}

			for _, bubble := range bubbles {
				if bubble.Y == y &&
					bubble.X >= 0 &&
					bubble.X < maxX {

					row[bubble.X] =
						ui.Bubble.Render(bubble.Art)
				}
			}

			fmt.Print(ui.Border.Render("│"))

			for _, cell := range row {
				fmt.Print(cell)
			}

			fmt.Println(ui.Border.Render("│"))
		}

		fmt.Println(
			ui.Border.Render(
				"└" + strings.Repeat("─", maxX) + "┘",
			),
		)

		fmt.Println(musicStatus)

		time.Sleep(400 * time.Millisecond)
	}
}
