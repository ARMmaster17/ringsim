package main

import "fmt"

import "math/rand"

import "os"

type fighter struct {
	name     string
	health   float32
	strength float32
	wins     int
	fights   int
	dead     bool
}

func main() {
	fmt.Print("Initializing fighters\n")
	var fighters [100]fighter
	for i := 0; i < 100; i++ {
		fighters[i] = fighter{name: generateFighterName(), health: 100, strength: float32(rand.Intn(80) + 1)}
	}
	fmt.Print("Fighters Initialized\n")
	for fight := 1; fight <= 1000; fight++ {
		fmt.Printf("Round %v\n", fight)
		p1 := &fighters[rand.Intn(len(fighters))]
		for {
			if !p1.dead {
				break
			}
			p1 = &fighters[rand.Intn(len(fighters))]
		}
		p2 := &fighters[rand.Intn(len(fighters))]
		for {
			if !p2.dead {
				break
			}
			p2 = &fighters[rand.Intn(len(fighters))]
		}
		fmt.Printf("%v VS %v\n", p1.name, p2.name)
		p1Damage := float32(rand.Intn(int(p1.strength)))
		fmt.Printf("%v deals %vhp damage\n", p1.name, p1Damage)
		if p1Damage >= p2.health {
			fmt.Printf("%v is out!\n", p2.name)
			p2.dead = true
		} else {
			p2.health = p2.health - p1Damage
			p2Damage := float32(rand.Intn(int(p2.strength)))
			fmt.Printf("%v deals %vhp damage\n", p2.name, p2Damage)
			if p2Damage >= p1.health {
				fmt.Printf("%v is out!\n", p1.name)
				p1.dead = true
			} else {
				p1.health = p1.health - p2Damage
			}
		}
		aliveCounter := 0
		for i := 0; i < len(fighters); i++ {
			if !fighters[i].dead {
				aliveCounter++
			}
		}
		fmt.Printf("%v fighters still alive.\n", aliveCounter)
		if aliveCounter == 1 {
			fmt.Print("Fight over!\n")
			for i := 0; i < len(fighters); i++ {
				if !fighters[i].dead {
					fmt.Printf("%v wins!\n", fighters[i].name)
					break
				}
			}
			os.Exit(0)
		} else if aliveCounter == 0 {
			fmt.Print("No winner\n")
			os.Exit(0)
		}
	}
	fmt.Print("Fight took too long. Everyone went home!\n")
	os.Exit(0)
}

func generateFighterName() string {
	return randSeq(6)
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
