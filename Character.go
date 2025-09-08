package main

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
+