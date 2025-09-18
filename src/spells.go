package main

import "fmt"

// Gestion des sorts
func useSpell(p *Character, m *Monster) {
	if len(p.Skills) == 0 {
		fmt.Println("Vous ne connaissez aucun sort.")
		return
	}

	fmt.Println("Sorts disponibles :")
	for i, s := range p.Skills {
		fmt.Printf("%d. %s\n", i+1, s)
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisir un sort : ")
	ch := readChoice()

	if ch == "0" {
		return
	}

	idx := parseIndex(ch)
	if idx < 1 || idx > len(p.Skills) {
		fmt.Println("Choix invalide.")
		return
	}

	spell := p.Skills[idx-1]
	switch spell {
	case "Coup de poing":
		// Sort de base : 8 dégâts, sans coût en mana
		dmg := 8
		m.HP -= dmg
		if m.HP < 0 {
			m.HP = 0
		}
		fmt.Printf("%s utilise %s et inflige %d dégâts à %s\n", p.Name, spell, dmg, m.Name)
		fmt.Printf("%s PV: %d/%d\n", m.Name, m.HP, m.HPMax)

	case "Boule de feu":
		// 18 dégâts, coûte 20 mana
		cost := 20
		if p.Mana < cost {
			fmt.Println("Mana insuffisant.")
			return
		}
		p.Mana -= cost
		dmg := 18
		m.HP -= dmg
		if m.HP < 0 {
			m.HP = 0
		}
		fmt.Printf("%s lance %s et inflige %d dégâts à %s (Mana restant %d/%d)\n",
			p.Name, spell, dmg, m.Name, p.Mana, p.ManaMax)
		fmt.Printf("%s PV: %d/%d\n", m.Name, m.HP, m.HPMax)

	default:
		fmt.Println("Sort inconnu.")
	}
}
