package main

import (
	"fmt"
	"math/rand"
	"time"
)

const timeMatch = 5
const maxLevel = 30
const creepsForNewLevel = 300

func denayCreeps(currCreeps int) (int, bool) {
	var newLevel bool
	var creepsNow int
	CreepsDenayd := rand.Intn(500)
	currCreeps += CreepsDenayd
	if currCreeps >= creepsForNewLevel {
		creepsNow = currCreeps - creepsForNewLevel
		newLevel = true
	} else {
		creepsNow = currCreeps
		newLevel = false
	}
	return creepsNow, newLevel
}

func upgradeLevel(currLevel uint, currCreeps int) (uint, int) {
	var upgrade bool
	currCreeps, upgrade = denayCreeps(currCreeps)
	if currLevel <= maxLevel-1 && upgrade == true {
		time.Sleep(30 * time.Millisecond)
		currLevel += 1
	}
	return currLevel, currCreeps
}

func CreateHero(name string) Heroes {
	return Heroes{name: name, level: 1}
}

type Heroes struct {
	name   string
	level  uint
	creeps int
}

func main() {

	listH := make([]Heroes, 0, 10)

	listH = append(listH,
		CreateHero("Axe"),
		CreateHero("Io"),
		CreateHero("Marci"),
		CreateHero("Pudge"),
		CreateHero("Doom"),
		CreateHero("Tusk"),
		CreateHero("Tiny"),
		CreateHero("Bane"),
		CreateHero("Sven"),
		CreateHero("Slark"),
	)
	var winner, looser string
	to := time.After(timeMatch * time.Second)
	done := make(chan bool, 1)

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(2)
	y := rand.Intn(9)

	if x == 0 {
		winner = "Radiant"
		looser = "Dire"
	} else {
		winner = "Dire"
		looser = "Radiant"
	}

	fmt.Println("Начало матча")
	go func() {
		defer fmt.Println("ГГ проебали", looser)
		for {
			select {
			case <-to:
				fmt.Println("Победили", winner)
				done <- true
				return
			default:
				y = rand.Intn(10)
				CurrHero := &listH[y]
				CurrHero.level, CurrHero.creeps = upgradeLevel(CurrHero.level, CurrHero.creeps)
			}
		}
	}()

	<-done
	fmt.Println("Конец матча")
	for i := 0; i < len(listH); i++ {
		CurrHero := listH[i]
		fmt.Println(CurrHero.name, "have", CurrHero.level, "level")
	}

}
