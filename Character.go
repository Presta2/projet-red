package main

import (
	"fmt"
	"os"
)

// Définition de la structure Character
type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []string
}

// Constructeur pour créer un nouveau personnage
func NewCharacter(name, class string, level, maxHP, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

// Méthode pour afficher les infos du personnage
func (c Character) DisplayInfo() {
	fmt.Printf("Nom       : %s\n", c.Name)
	fmt.Printf("Classe    : %s\n", c.Class)
	fmt.Printf("Niveau    : %d\n", c.Level)
	fmt.Printf("PV        : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Inventaire: %v\n", c.Inventory)
}

// Méthode pour afficher l'inventaire
func (c Character) AccessInventory() {
	fmt.Println("Inventaire du personnage :")
	if len(c.Inventory) == 0 {
		fmt.Println("L’inventaire est vide.")
	} else {
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
}

// Fonction principale
func main() {
	// Création d'un personnage
	char := NewCharacter("Arthas", "Guerrier", 5, 100, 85, []string{"Épée", "Bouclier", "Potion de soin"})

	// Boucle du menu
	for {
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l’inventaire")
		fmt.Println("3. Quitter")
		fmt.Print("Choisis une option : ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erreur de saisie, réessaie...")
			continue
		}

		switch choice {
		case 1:
			char.DisplayInfo()
		case 2:
			char.AccessInventory()
		case 3:
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
