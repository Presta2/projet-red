package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

// Lire une entrée utilisateur
func readChoice() string {
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}

// Convertir une entrée en int
func parseIndex(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}

// Pause avec confirmation utilisateur
func pause() {
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	_, _ = reader.ReadString('\n')
}

// Petite pause automatique
func pauseShort() {
	time.Sleep(600 * time.Millisecond)
}

// Effacer l'écran
func clearScreen() {
	fmt.Print("\n\n")
}

// Vérifier si une valeur est dans un slice
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Compter le nombre d'occurrences d'une valeur dans un slice
func countInSlice(slice []string, value string) int {
	cnt := 0
	for _, v := range slice {
		if v == value {
			cnt++
		}
	}
	return cnt
}

// Supprimer la première occurrence d'une valeur dans un slice
func removeFromSlice(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Vérifie que le nom ne contient que des lettres (simplifié)
func isAlpha(s string) bool {
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			continue
		}
		// on tolère les espaces
		if r == ' ' {
			continue
		}
		return false
	}
	return true
}
