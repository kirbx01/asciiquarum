package main

import ("fmt"
		"time")
type Fish struct {
	Name string
	X int
	Y int
	speed int
}

func main(){
	fmt.Println("Welcome to my aquarium....Bleep...Bloop...Bloop 𓆝 ⋆｡𖦹°‧ ")

	fish := Fish{
    Name:  "Marias",
    X:     10,
    Y:     5,
    speed: 1,
}

for i := 0; i < 5; i++ {
	fish.X = fish.X + 2*fish.speed

	fmt.Printf("Fish: %s, Position: (%d, %d), Speed: %d\n", fish.Name, fish.X, fish.Y, fish.speed)

	time.Sleep(500 * time.Millisecond)
}
}

