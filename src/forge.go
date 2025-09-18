package main

import "fmt"

func ForgeMenu(p *Character) {
	for {
		clearScreen()
		fmt.Println("=== FORGERON ===")
		fmt.Printf("Or: %d\n", p.Gold)
		fmt.Println("Objets à fabriquer (coût de fabrication : 5 pièces d'or en plus des matériaux) :")

		i := 1
		order := []string{
			"Chapeau de l'aventurier",
			"Tunique de l'aventurier",
			"Bottes de l'aventurier",
		}
		for _, name := range order {
			fmt.Printf("%d. %s\n", i, name)
			i++
		}
		fmt.Println("0. Retour")
		fmt.Print("Choix : ")
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
		costMoney := 5

		if p.Gold < costMoney {
			fmt.Println("Pas assez d'or pour payer la fabrication.")
			pause()
			continue
		}

		// Vérification des ressources nécessaires
		recipe := CraftRecipes[item]
		can := true
		for mat, q := range recipe {
			if count := countInSlice(p.Inventory, mat); count < q {
				fmt.Printf("Il vous manque %d x %s\n", q-count, mat)
				can = false
			}
		}
		if !can {
			pause()
			continue
		}

		// Vérification de la place dans l'inventaire
		materialsRemoved := 0
		for _, q := range recipe {
			materialsRemoved += q
		}
		newLen := len(p.Inventory) - materialsRemoved + 1
		if newLen > p.InventoryLimit {
			fmt.Println("Pas assez de place dans l'inventaire pour l'objet fabriqué.")
			pause()
			continue
		}

		// Retirer les matériaux
		for mat, q := range recipe {
			for j := 0; j < q; j++ {
				p.removeOneFromInventory(mat)
			}
		}

		// Déduire le coût et ajouter l’objet
		p.Gold -= costMoney
		p.Inventory = append(p.Inventory, item)
		fmt.Printf("Fabrication de %s réussie. Vous payez %d pièces.\n", item, costMoney)
		pauseShort()
	}
}
