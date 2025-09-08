package main

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []string
}

func initCharacter(name, class string, level, maxHP, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

func main()

func (c Character) displayInfo() {
	fmt.Printf("Nom : %s\n", c.Name)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Inventaire : %v\n", c.Inventory)
}
