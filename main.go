package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	player1Distance = 0
	player2Distance = 0
	rubinho         = "Rubinho"
	lock            = sync.Mutex{}
	limit           = 0
)

func runPlayer1(name string) {
	player1Distance = player1Distance + 1

	if player1Distance == 250 {

		lock.Lock()
		rubinho = "Rubinho arrived!"
		lock.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player1Distance)

		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	} else if player1Distance == 500 {

		lock.Lock()
		rubinho = "Rubinho is ready to go!"
		lock.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player1Distance)

		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	} else if player1Distance == 850 {

		lock.Lock()
		rubinho = "Has anyone seen Rubinho?"
		lock.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player1Distance)

		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	}

}

func runPlayer2(name string) {
	player2Distance = player2Distance + 1

	if player2Distance == 250 {

		lock.Lock()
		rubinho = "Rubinho is very excited!"
		lock.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player2Distance)
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	} else if player2Distance == 500 {

		lock.Lock()
		rubinho = "Rubinho is ready to go"
		lock.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player2Distance)
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)

	} else if player2Distance == 850 {

		lock.Lock()
		rubinho = "Where is Rubinho?"
		lock.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player2Distance)
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	}

}

func printFunc(p1, p2 string) {
	if player1Distance < player2Distance {
		fmt.Printf("%s first! %d Kms\n", p1, player1Distance)
	} else {
		fmt.Printf("%s first! %d Kms\n", p2, player2Distance)
	}
}

func main() {
	for {
		limit = limit + 1

		if limit > 1000 {
			fmt.Println("Race finished!")
			break
		}
		go runPlayer1("Bolt")
		go runPlayer2("Flash")

		printFunc("Bolt", "Flash")

	}

	time.Sleep(2 * time.Second)
	fmt.Printf("The winner is...")
	time.Sleep(3 * time.Second)
	printFunc("Bolt", "Flash")
	fmt.Println("Rubinho started!!!")
}
