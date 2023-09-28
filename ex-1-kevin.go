// package main

// import "fmt"

// // =======
// // ÉNONCÉ
// // =======

// // Créer une fonction ft_strstr qui repère la 1ère occurrence
// // d'une sous-chaine dans la chaine principale
// // Et retourne l'index de celle-ci dans la chaine principale
// // (Si aucune occurrence : retourne nil)

// func main() {

// 	firstString := "Hello World!"
// 	secondString := "World"

// 	result := ft_strstr(firstString, secondString)
// 	fmt.Println(result)

// }

// func ft_strstr(s1 string, s2 string) int {

// 	// J'initialise l'index à -1
// 	index := -1

// 	// Je repère la 1e lettre de la 2e chaine de caractères
// 	// (pas obligatoire mais je trouve cela plus lisible)
// 	firstLetterOfSecondString := s2[0]

// 	if len(s2) > len(s1) {
// 		return -1
// 	}

// 	// Je parcours chaque lettre de la 1ère chaine de caractères
// 	for i := 0; i < len(s1); i++ {
// 		// Si la lettre est égale à celle que je cherche
// 		if s1[i] == firstLetterOfSecondString {
// 			index = i
// 			// Je parcours toute la 2e chaine en comparant les lettres
// 			for j := 0; j < len(s2); j++ {
// 				if s1[i+j] != s2[j] {
// 					index = -1
// 				}
// 			}
// 			break
// 		}
// 	}
// 	return index
// }
