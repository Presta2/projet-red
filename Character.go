package main

import "fmt"

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
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

// Fonction principale
func main() {
	// Création d'un personnage
	char := NewCharacter("Arthas", "Guerrier", 5, 100, 85, []string{"Épée", "Bouclier", "Potion de soin"})

	// Affichage des infos et de l'inventaire
	char.DisplayInfo()
	fmt.Println()
	char.AccessInventory()
}
