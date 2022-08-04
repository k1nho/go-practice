package main

import (
	"fmt"
	"sync"
)

func main() {
	// in order to synchronize routines we can use waitgroups
	var signal sync.WaitGroup
	signal.Add(2) // number of routines that need to be completed in order for the main function to finish
	go invokeCard("Dark Magician", &signal)
	go invokeSpell("Toon Field", &signal)
	signal.Wait() // we wait for the signal to be done
	fmt.Println("Turn ends\n")
}

func invokeCard(card string, signal *sync.WaitGroup) {
	fmt.Printf("Come forth %v\n", card)
	signal.Done()
}

func invokeSpell(spellCard string, signal *sync.WaitGroup) {
	fmt.Printf("Active the spell card %v \n", spellCard)
	signal.Done()
}
