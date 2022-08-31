package main

import (
	"fmt"
	"sync"
	"time"
)

type rubinhoMutex struct {
	sync.Mutex
	v string
}

var (
	player1Distance = 0
	player2Distance = 0
	rubinho         = rubinhoMutex{v: ""}
	lock            = sync.Mutex{}
	limit           = 0
	overtaking      = 0
)

func runPlayer1(name string) {

	if player1Distance == player2Distance {
		lock.Lock()
		overtaking += 1
		lock.Unlock()
	}
	player1Distance += 1

	if player1Distance == 250 {

		rubinho.Lock()
		rubinho.v = "Rubinho arrived!"
		rubinho.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player1Distance)

		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	} else if player1Distance == 500 {

		rubinho.Lock()
		rubinho.v = "Rubinho is ready to go!"
		rubinho.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player1Distance)

		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	} else if player1Distance == 850 {

		rubinho.Lock()
		rubinho.v = "Has anyone seen Rubinho?"
		rubinho.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player1Distance)

		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	}

}

func runPlayer2(name string) {

	if player1Distance == player2Distance {
		lock.Lock()
		overtaking += 1
		lock.Unlock()
	}
	player2Distance += 1

	if player2Distance == 250 {

		rubinho.Lock()
		rubinho.v = "Rubinho is very excited!"
		rubinho.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player2Distance)
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	} else if player2Distance == 500 {

		rubinho.Lock()
		rubinho.v = "Rubinho is ready to go"
		rubinho.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player2Distance)
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)

	} else if player2Distance == 850 {

		rubinho.Lock()
		rubinho.v = "Where is Rubinho?"
		rubinho.Unlock()

		fmt.Printf("%s: %d Kms\n", name, player2Distance)
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%s\n", rubinho)
	}

}

func announcer(p1, p2 string) {
	if player1Distance > player2Distance {
		fmt.Printf("In %d, %s first! %d Kms and %s second with %d Kms\n", limit, p1, player1Distance, p2, player2Distance)
	} else if player1Distance < player2Distance {
		fmt.Printf("In %d, %s first! %d Kms and %s second with %d Kms\n", limit, p2, player2Distance, p1, player1Distance)
	} else {
		fmt.Printf("In %d, %s and %s with the same Kms, %d, %d \n", limit, p1, p2, player1Distance, player2Distance)
	}
}

func printFunc(p1, p2 string) {
	if player1Distance > player2Distance {
		fmt.Printf("%s first! %d Kms\n", p1, player1Distance)
	} else {
		fmt.Printf("%s first! %d Kms\n", p2, player2Distance)
	}
}

func main() {
	for {
		limit = limit + 1

		if limit > 1000 {
			break
		}
		go runPlayer1("Bolt")
		go runPlayer2("Flash")
		go announcer("Bolt", "Flash")

		printFunc("Bolt", "Flash")

	}

	time.Sleep(2 * time.Second)
	fmt.Println("Race finished!")
	time.Sleep(1 * time.Second)
	fmt.Printf("Overtaking %d times\n", overtaking)
	fmt.Printf("The winner is...\n")
	time.Sleep(2 * time.Second)
	printFunc("Bolt", "Flash")
	fmt.Println("Rubinho started!!!")
}
