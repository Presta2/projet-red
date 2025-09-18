package main

// Liste des objets disponibles chez le marchand
// et leur prix en pièces d'or
var ShopItems = map[string]int{
	"Potion de vie":                3,
	"Potion de poison":             6,
	"Livre de Sort : Boule de feu": 25,
	"Potion de mana":               5,
	"Fourrure de Loup":             4,
	"Peau de Troll":                7,
	"Cuir de Sanglier":             3,
	"Plume de Corbeau":             1,
	"Augmentation d'inventaire":    30,
}

// Recettes de craft pour le forgeron
var CraftRecipes = map[string]map[string]int{
	"Chapeau de l'aventurier": {
		"Plume de Corbeau": 1,
		"Cuir de Sanglier": 1,
	},
	"Tunique de l'aventurier": {
		"Fourrure de Loup": 2,
		"Peau de Troll":    1,
	},
	"Bottes de l'aventurier": {
		"Fourrure de Loup": 1,
		"Cuir de Sanglier": 1,
	},
}
