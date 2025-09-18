package main

import (
	"fmt"
)

// Combat d'entraînement
func TrainingFight(p *Character) {
	clearScreen()
	fmt.Println("=== COMBAT D'ENTRAÎNEMENT ===")
	g := InitGoblin()
	turn := 1

	// initiative: qui commence
	if p.Initiative >= g.Initiative {
		fmt.Println("Vous commencez !")
	} else {
		fmt.Println("Le gobelin commence...")
	}

	for {
		fmt.Printf("\n--- Tour %d ---\n", turn)

		if p.Initiative >= g.Initiative {
			// joueur commence
			characterTurn(p, &g)
			if g.IsDead() {
				fmt.Println("Vous avez vaincu le Gobelin !")
				grantXP(p, 10)
				return
			}
			goblinPattern(p, &g, turn)
			if p.IsDead() {
				fmt.Println("Vous avez été vaincu...")
				p.Resurrect()
				return
			}
		} else {
			// gobelin commence
			goblinPattern(p, &g, turn)
			if p.IsDead() {
				fmt.Println("Vous avez été vaincu...")
				p.Resurrect()
				return
			}
			characterTurn(p, &g)
			if g.IsDead() {
				fmt.Println("Vous avez vaincu le Gobelin !")
				grantXP(p, 10)
				return
			}
		}
		turn++
	}
}

// Tour du joueur
func characterTurn(p *Character, m *Monster) {
	for {
		fmt.Println("\nVotre tour :")
		fmt.Println("1. Attaquer (Coup de poing - 5 dégâts)")
		fmt.Println("2. Utiliser Inventaire")
		fmt.Println("3. Utiliser Sort")
		fmt.Println("0. Passer")
		fmt.Print("Choix : ")
		ch := readChoice()

		switch ch {
		case "1":
			dmg := 5
			m.HP -= dmg
			if m.HP < 0 {
				m.HP = 0
			}
			fmt.Printf("%s inflige %d dégâts à %s\n", p.Name, dmg, m.Name)
			fmt.Printf("%s PV: %d/%d\n", m.Name, m.HP, m.HPMax)
			return
		case "2":
			p.AccessInventory()
			return
		case "3":
			useSpell(p, m)
			return
		case "0":
			fmt.Println("Vous passez votre tour.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Tour du gobelin
func goblinPattern(p *Character, g *Monster, turn int) {
	mult := 1
	if turn%3 == 0 {
		mult = 2
	}
	dmg := g.Attack * mult
	p.HP -= dmg
	if p.HP < 0 {
		p.HP = 0
	}
	fmt.Printf("%s inflige %d dégâts à %s\n", g.Name, dmg, p.Name)
	fmt.Printf("%s PV: %d/%d\n", p.Name, p.HP, p.HPMax)
}
