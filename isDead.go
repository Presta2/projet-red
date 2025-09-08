package main

import (
	"fmt"
)

type Player struct {
	Name    string
	HP      int
	MaxHP   int
	IsAlive bool
}

func isDead(p *Player) {
	if p.HP <= 0 {
		fmt.Printf("%s est mort.\n", p.Name)
		p.IsAlive = false

		p.HP = p.MaxHP / 2
		p.IsAlive = true
		fmt.Printf("%s est ressuscité avec %d points de vie.\n", p.Name, p.HP)
	} else {
		fmt.Printf("%s est toujours en vie avec %d points de vie.\n", p.Name, p.HP)
	}
}

func Joueur() {
	player := Player{Name: "Arthas", HP: 0, MaxHP: 100, IsAlive: true}
	isDead(&player)
}
