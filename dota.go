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
	currCreeps = currCreeps + CreepsDenayd
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
		currLevel = currLevel + 1
	}
	return currLevel, currCreeps
}

func main() {
	type Heroes struct {
		name   string
		level  uint
		creeps int
	}
	listH := make([]Heroes, 0)

	var CurrHero Heroes

	var Axe Heroes
	Axe.name = "Axe"
	Axe.level = 1
	Axe.creeps = 0
	listH = append(listH, Axe)

	var Io Heroes
	Io.name = "Io"
	Io.level = 1
	Io.creeps = 0
	listH = append(listH, Io)

	var Slark Heroes
	Slark.name = "Slark"
	Slark.level = 1
	Slark.creeps = 0
	listH = append(listH, Slark)

	var Marci Heroes
	Marci.name = "Marci"
	Marci.level = 1
	Marci.creeps = 0
	listH = append(listH, Marci)

	var Pudge Heroes
	Pudge.name = "Pudge"
	Pudge.level = 1
	Pudge.creeps = 0
	listH = append(listH, Pudge)

	var Doom Heroes
	Doom.name = "Doom"
	Doom.level = 1
	Doom.creeps = 0
	listH = append(listH, Doom)

	var Sven Heroes
	Sven.name = "Sven"
	Sven.level = 1
	Sven.creeps = 0
	listH = append(listH, Sven)

	var Tusk Heroes
	Tusk.name = "Tusk"
	Tusk.level = 1
	Tusk.creeps = 0
	listH = append(listH, Tusk)

	var Tiny Heroes
	Tiny.name = "Tiny"
	Tiny.level = 1
	Tiny.creeps = 0
	listH = append(listH, Tiny)

	var Bane Heroes
	Bane.name = "Bane"
	Bane.level = 1
	Bane.creeps = 0
	listH = append(listH, Bane)

	var winner string
	to := time.After(timeMatch * time.Second)
	done := make(chan bool, 1)

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(2)
	y := rand.Intn(9)

	if x == 0 {
		winner = "Radiant"
	} else {
		winner = "Dire"
	}

	fmt.Println("Начало матча")
	go func() {
		defer fmt.Println("ГГ проебали")
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
		CurrHero = listH[i]
		fmt.Println(CurrHero.name, "have", CurrHero.level, "level")
	}

}
