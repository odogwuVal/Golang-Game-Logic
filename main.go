package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	max_usage = 5
	usage     = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Want to try?`
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, max_usage)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for i := 0; i < max_usage; i++ {
		n := rand.Intn(guess + 1)
		switch {
		case i == 0 && n == guess:
			fmt.Println("Hurray!! first turn winner")
			return
		case n == guess:
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
