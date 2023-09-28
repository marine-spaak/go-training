package main

import (
	"fmt"
)

// =======
// ÉNONCÉ
// =======

// Créer une liste chainée contenant un int comme valeur. Lui ajouter les méthodes suivantes
// Ajout d'un noeud (donc nouvelle valeur)
// Suppression d'un noeud (en fonction de sa valeur)
// Obtenir le premier noeud
// Obtenir le dernier noeud

func main() {

	myString := "Hello;World;!"

	result := ft_split(myString, ";")
	fmt.Println(result)

}

func ft_split(stringToSplit string, separator string) []string {

	// J'initialise le résultat
	var resultArray []string
	substring := ""

	// Je parcours la chaine à séparer
	for i := 0; i < len(stringToSplit); i++ {
		// Si je tombe sur une lettre normale (pas le séparateur)
		if string(stringToSplit[i]) != separator {
			// Alors je l'ajoute à la sous-chaine que je suis en train de construire
			substring = substring + string(stringToSplit[i])

			// Sinon (donc si je tombe sur le séparateur)
		} else {
			// J'ajoute la chaine actuelle au tableau de résultats
			resultArray = append(resultArray, substring)
			// Je remets la sous-chaine à zéro
			substring = ""
		}
	}
	// Il faut que j'ajoute le dernier bout de chaine (entre le dernier séparateur et la fin)
	resultArray = append(resultArray, substring)

	// Je renvoie le tableau complet
	return resultArray
}
