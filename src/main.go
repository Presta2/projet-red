package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bienvenue dans Projet RED - version Go")
	player := Character{}
	characterCreation(&player)

	for {
		clearScreen()
		fmt.Println("=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Entraînement (Combat)")
		fmt.Println("6. Augmenter l'inventaire (30 pièces, max 3 fois)")
		fmt.Println("7. Quitter")

		choice := readChoice()
		switch choice {
		case "1":
			player.DisplayInfo()
			pause()
		case "2":
			player.AccessInventory()
		case "3":
			ShopMenu(&player)
		case "4":
			ForgeMenu(&player)
		case "5":
			TrainingFight(&player)
			pause()
		case "6":
			if player.InventoryUpgrades >= 3 {
				fmt.Println("Vous avez déjà utilisé les 3 upgrades disponibles.")
			} else {
				if player.Gold < 30 {
					fmt.Println("Pas assez d'or.")
				} else {
					if player.UpgradeInventorySlot() {
						player.Gold -= 30
						fmt.Println("Capacité d'inventaire augmentée !")
					}
				}
			}
			pause()
		case "7":
			fmt.Println("Au revoir aventurier !")
			return
		default:
			fmt.Println("Choix invalide.")
			pause()
		}
	}
}
