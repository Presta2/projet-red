package main

import "fmt"

// Menu du marchand
func ShopMenu(p *Character) {
	for {
		clearScreen()
		fmt.Println("=== MARCHAND ===")
		fmt.Printf("Or: %d\n", p.Gold)

		order := []string{
			"Potion de vie",
			"Potion de poison",
			"Livre de Sort : Boule de feu",
			"Potion de mana",
			"Fourrure de Loup",
			"Peau de Troll",
			"Cuir de Sanglier",
			"Plume de Corbeau",
			"Augmentation d'inventaire",
		}

		for i, name := range order {
			if cost, ok := ShopItems[name]; ok {
				fmt.Printf("%d. %s (%d pièces)\n", i+1, name, cost)
			}
		}

		fmt.Println("0. Retour")
		fmt.Print("Choisissez un article à acheter (numéro) : ")
		choice := readChoice()

		if choice == "0" {
			return
		}

		idx := parseIndex(choice)
		if idx < 1 || idx > len(order) {
			fmt.Println("Choix invalide.")
			pause()
			continue
		}

		item := order[idx-1]
		cost := ShopItems[item]

		// Cas particulier : upgrade inventaire
		if item == "Augmentation d'inventaire" {
			if p.InventoryUpgrades >= 3 {
				fmt.Println("Vous avez déjà utilisé les 3 upgrades.")
				pause()
				continue
			}
			if p.Gold < cost {
				fmt.Println("Pas assez d'or.")
				pause()
				continue
			}
			p.Gold -= cost
			p.UpgradeInventorySlot()
			fmt.Println("Inventaire augmenté !")
			pause()
			continue
		}

		// Vérifier l'or
		if p.Gold < cost {
			fmt.Println("Pas assez d'or.")
			pause()
			continue
		}

		// Vérifier la place dans l'inventaire
		if len(p.Inventory) >= p.InventoryLimit {
			fmt.Println("Inventaire plein, impossible d'acheter.")
			pause()
			continue
		}

		// Achat réussi
		p.Inventory = append(p.Inventory, item)
		p.Gold -= cost
		fmt.Printf("Vous achetez %s pour %d pièces.\n", item, cost)
		pause()
	}
}
