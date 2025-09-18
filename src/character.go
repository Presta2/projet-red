package main

import (
	"fmt"
	"strings"
	"time"
)

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name              string
	Class             string
	Level             int
	HPMax             int
	HP                int
	Inventory         []string
	InventoryLimit    int
	InventoryUpgrades int
	Gold              int
	Skills            []string
	Equipment         Equipment
	ManaMax           int
	Mana              int
	Experience        int
	ExperienceMax     int
	Initiative        int
}

func (c *Character) DisplayInfo() {
	fmt.Println("=== FICHE PERSONNAGE ===")
	fmt.Printf("Nom: %s\n", c.Name)
	fmt.Printf("Classe: %s\n", c.Class)
	fmt.Printf("Niveau: %d\n", c.Level)
	fmt.Printf("PV: %d / %d\n", c.HP, c.HPMax)
	fmt.Printf("Mana: %d / %d\n", c.Mana, c.ManaMax)
	fmt.Printf("Or: %d\n", c.Gold)
	fmt.Printf("Inventaire (%d/%d): %v\n", len(c.Inventory), c.InventoryLimit, c.Inventory)
	fmt.Printf("Compétences: %v\n", c.Skills)
	fmt.Printf("Equipement: Tete=%s Torse=%s Pieds=%s\n", c.Equipment.Head, c.Equipment.Torso, c.Equipment.Feet)
	fmt.Printf("XP: %d / %d\n", c.Experience, c.ExperienceMax)
	fmt.Printf("Initiative: %d\n", c.Initiative)
}

func (c *Character) AccessInventory() {
	for {
		clearScreen()
		fmt.Println("=== INVENTAIRE ===")
		if len(c.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
			pause()
			return
		}
		for i, it := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, it)
		}
		fmt.Println("0. Retour")
		fmt.Print("Choisissez un objet à utiliser (numéro) : ")
		choice := readChoice()
		if choice == "0" {
			return
		}
		idx := parseIndex(choice)
		if idx < 1 || idx > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			pause()
			continue
		}
		item := c.Inventory[idx-1]
		c.UseItem(item)
		pause()
	}
}

func (c *Character) UseItem(item string) {
	switch item {
	case "Potion de vie":
		if c.HP >= c.HPMax {
			fmt.Println("Vos PV sont déjà au maximum.")
			return
		}
		c.HP += 50
		if c.HP > c.HPMax {
			c.HP = c.HPMax
		}
		c.removeOneFromInventory(item)
		fmt.Printf("Vous utilisez %s. PV: %d/%d\n", item, c.HP, c.HPMax)
	case "Potion de poison":
		c.removeOneFromInventory(item)
		fmt.Println("Vous avez bu une potion de poison !")
		poisonPot(c)
	case "Livre de Sort : Boule de feu":
		c.removeOneFromInventory(item)
		if contains(c.Skills, "Boule de feu") {
			fmt.Println("Vous connaissez déjà Boule de feu.")
		} else {
			c.Skills = append(c.Skills, "Boule de feu")
			fmt.Println("Vous apprenez 'Boule de feu' !")
		}
	case "Potion de mana":
		if c.Mana >= c.ManaMax {
			fmt.Println("Votre mana est déjà au maximum.")
			return
		}
		c.Mana += 30
		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}
		c.removeOneFromInventory(item)
		fmt.Printf("Vous utilisez %s. Mana: %d/%d\n", item, c.Mana, c.ManaMax)
	case "Chapeau de l'aventurier":
		c.equipItem("head", item)
	case "Tunique de l'aventurier":
		c.equipItem("torso", item)
	case "Bottes de l'aventurier":
		c.equipItem("feet", item)
	default:
		fmt.Printf("L'utilisation de %s n'est pas encore implémentée.\n", item)
	}
}

func (c *Character) removeOneFromInventory(item string) {
	for i, it := range c.Inventory {
		if it == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
}

func (c *Character) equipItem(slot, item string) {
	var replaced string
	switch slot {
	case "head":
		replaced = c.Equipment.Head
		c.Equipment.Head = item
		c.adjustHPForEquip(item, replaced)
	case "torso":
		replaced = c.Equipment.Torso
		c.Equipment.Torso = item
		c.adjustHPForEquip(item, replaced)
	case "feet":
		replaced = c.Equipment.Feet
		c.Equipment.Feet = item
		c.adjustHPForEquip(item, replaced)
	}
	// retirer l’objet équipé de l’inventaire, remettre l’ancien si besoin
	c.removeOneFromInventory(item)
	if replaced != "" {
		if len(c.Inventory) < c.InventoryLimit {
			c.Inventory = append(c.Inventory, replaced)
			fmt.Printf("Vous remplacez %s par %s. %s retourne dans l'inventaire.\n", replaced, item, replaced)
		} else {
			fmt.Printf("Vous remplacez %s par %s. Mais votre inventaire est plein, l'ancien équipement est perdu.\n", replaced, item)
		}
	} else {
		fmt.Printf("Vous équipez %s.\n", item)
	}
}

func (c *Character) adjustHPForEquip(newItem, oldItem string) {
	// appliquer bonus/malus
	apply := func(it string, sign int) {
		switch it {
		case "Chapeau de l'aventurier":
			c.HPMax += 10 * sign
		case "Tunique de l'aventurier":
			c.HPMax += 25 * sign
		case "Bottes de l'aventurier":
			c.HPMax += 15 * sign
		}
	}
	if oldItem != "" {
		apply(oldItem, -1)
	}
	if newItem != "" {
		apply(newItem, +1)
	}
	if c.HP > c.HPMax {
		c.HP = c.HPMax
	}
}

func (c *Character) IsDead() bool {
	return c.HP <= 0
}

func (c *Character) Resurrect() {
	fmt.Println("Vous êtes mort... Résurrection automatique !")
	c.HP = c.HPMax / 2
	fmt.Printf("Vous revenez à %d/%d PV\n", c.HP, c.HPMax)
}

func (c *Character) UpgradeInventorySlot() bool {
	if c.InventoryUpgrades >= 3 {
		return false
	}
	c.InventoryLimit += 10
	c.InventoryUpgrades++
	return true
}

// personnage

func characterCreation(c *Character) {
	for {
		fmt.Print("Entrez le nom de votre personnage (lettres seulement) : ")
		name := readChoice()
		name = strings.TrimSpace(name)
		if name == "" || !isAlpha(name) {
			fmt.Println("Nom invalide. Utilisez uniquement des lettres.")
			continue
		}
		// formatage : première lettre majuscule
		name = strings.Title(strings.ToLower(name))
		c.Name = name
		break
	}

	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Humain (PV max 100)")
	fmt.Println("2. Elfe (PV max 80)")
	fmt.Println("3. Nain (PV max 120)")
	fmt.Print("Choix : ")
	choice := readChoice()
	switch choice {
	case "1":
		c.Class = "Humain"
		c.HPMax = 100
	case "2":
		c.Class = "Elfe"
		c.HPMax = 80
	case "3":
		c.Class = "Nain"
		c.HPMax = 120
	default:
		fmt.Println("Classe invalide, Humain par défaut.")
		c.Class = "Humain"
		c.HPMax = 100
	}
	c.Level = 1
	c.HP = c.HPMax / 2
	c.Inventory = []string{"Potion de vie", "Potion de vie", "Potion de vie"}
	c.InventoryLimit = 10
	c.InventoryUpgrades = 0
	c.Gold = 100
	c.Skills = []string{"Coup de poing"}
	c.Equipment = Equipment{}
	c.ManaMax = 50
	c.Mana = c.ManaMax
	c.Experience = 0
	c.ExperienceMax = 20
	c.Initiative = 10
}

func poisonPot(c *Character) {
	// inflige 10 PV par seconde pendant 3s
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("Poison: PV %d/%d\n", c.HP, c.HPMax)
		if c.IsDead() {
			c.Resurrect()
			return
		}
	}
}

func grantXP(c *Character, amount int) {
	fmt.Printf("%s gagne %d points d'expérience !\n", c.Name, amount)
	c.Experience += amount

	for c.Experience >= c.ExperienceMax {
		c.Experience -= c.ExperienceMax
		c.Level++
		c.ExperienceMax += 10
		c.HPMax += 10
		c.ManaMax += 5
		c.HP = c.HPMax
		c.Mana = c.ManaMax

		fmt.Printf(">>> LEVEL UP ! %s passe au niveau %d <<<\n", c.Name, c.Level)
		fmt.Printf("Nouveaux PV max : %d | Mana max : %d | XP pour prochain niveau : %d\n",
			c.HPMax, c.ManaMax, c.ExperienceMax)
	}
}
