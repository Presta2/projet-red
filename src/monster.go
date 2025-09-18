package main

// Structure d'un monstre
type Monster struct {
	Name       string
	HPMax      int
	HP         int
	Attack     int
	Initiative int
}

// Initialisation d'un Gobelin d'entraînement
func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entraînement",
		HPMax:      40,
		HP:         40,
		Attack:     5,
		Initiative: 5,
	}
}

// Vérifie si le monstre est mort
func (m *Monster) IsDead() bool {
	return m.HP <= 0
}
