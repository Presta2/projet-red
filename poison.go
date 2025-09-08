package main

import (
	"fmt"
)

func Poison() {
	inventaire := []string{}
	var choix int

	fmt.Println("Bienvenue chez le marchand !")
	fmt.Println("1 - Potion de vie")
	fmt.Println("2 - Potion de poison")
	fmt.Print("Que voulez-vous acheter ? ")
	fmt.Scanln(&choix)

	if choix == 1 {
		inventaire = append(inventaire, "Potion de vie")
	} else if choix == 2 {
		inventaire = append(inventaire, "Potion de poison")
	} else {
		fmt.Println("Choix invalide.")
	}

	fmt.Println("Inventaire :", inventaire)
}
