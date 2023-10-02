package main

import (
	"fmt"
)

// =======
// ÉNONCÉ
// =======

// Créer une liste chainée contenant un int comme valeur. Lui ajouter les méthodes suivantes
// Obtenir le premier noeud
// Obtenir le dernier noeud
// Ajout d'un noeud (donc nouvelle valeur) - pas 2 noeuds identiques
// Suppression d'un noeud (en fonction de sa valeur)

// ========================================
// DEFINITION DE "NODE" ET "LINKEDLIST"
// ========================================

// Définition de la structure de "node" (un noeud de la liste chaînée)
type node struct {
	data int   // Valeur du noeud (un entier)
	next *node // Pointeur vers le prochain noeud
}

// Description de la structure de "linkedList" (une liste chaînée)
type linkedList struct {
	head   *node // Pointeur vers le premier noeud
	length int   // Longueur de la liste
}

// =================================
// AFFICHER LES DONNEES DE LA LISTE
// =================================

// Remarque : ici le récepteur est "linkedList" et non pas "*linkedList"
// car on montre les données sans toucher à la structure de la liste

func (l linkedList) printListData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data) // Afficher un entier qui a pour valeur la donnée de toPrint
		toPrint = toPrint.next          // Au prochain tour, on affichera le noeud suivant
		l.length--                      // Je réduis la longueur qui tombera à 0 quand j'aurai parcouru toute la liste
	}
	fmt.Println("\n") // Saut de ligne
}

// ========================================
// TROUVER LE PREMIER OU LE DERNIER NOEUD
// ========================================

// Obtenir le premier noeud
func (l *linkedList) getFirstNode() *node {
	return l.head
}

// Obtenir le dernier noeud
func (l *linkedList) getLastNode() *node {
	current := l.head
	// Tant que je ne suis pas arrivée au bout, je récupère le suivant
	for current.next != nil {
		current = current.next
	}
	return current
}

// ==========================================
// TROUVER UN NOEUD EN FONCTION DE SA VALEUR
// ==========================================

func (l *linkedList) findNodeByValue(v int) *node {
	if l.head == nil {
		// CAS 1 - LISTE VIDE (résultat nil)
		return nil

	} else {
		// CAS 2 - LISTE NON VIDE
		// Initialiser une variable dans une boucle FOR si on ne s'en sert pas ailleurs
		for current := l.head; current.next != nil; current = current.next {
			if current.data == v { // Je check si la data du noeud vaut ce que je cherche
				return current // Si c'est le cas, je renvoie le noeud actuel
			}
		}
		// Si aucun n'a marché, je renvoie nil
		return nil
	}
}

// ==========================================
// AJOUTER UN NOEUD (SANS CRÉER DE DOUBLONS)
// ==========================================

// METHODE PREPREND (ajouter un noeud au début de la liste)
func (l *linkedList) prepend(n *node) {
	if l.head == nil {
		// CAS 1 - LISTE VIDE
		l.head = n
		l.length++
	} else {
		// CAS 2 - LISTE NON VIDE
		// Je vérifie que la liste ne contient pas encore cette valeur
		if l.findNodeByValue(n.data) == nil {
			second := l.head     // Création d'une variable "second" recevant la head actuelle de la liste
			l.head = n           // Le noeud reçu en argument devient la tête de la liste
			l.head.next = second // Le noeud suivant sera "second" (l'ancienne tête de liste)
			l.length++           // J'augmente la taille de la liste
		}
	}
}

// PREPEND (création d'un nouveau noeud à partir d'un entier)
func (l *linkedList) intPrepend(v int) {
	n := &node{} // Le & crée un pointeur
	n.data = v
	l.prepend(n)
}

// METHODE APPEND (ajouter un noeud en fin de liste)
func (l *linkedList) append(n *node) {
	if l.head == nil {
		// CAS 1 - LISTE VIDE
		l.head = n
		l.length++
	} else {
		// CAS 2 - LISTE NON VIDE
		// J'obtiens le dernier noeud et je lui ajoute un suivant
		last := l.getLastNode()
		last.next = n
		l.length++
	}
}

// APPEND (création d'un nouveau noeud à partir d'un entier)
func (l *linkedList) intAppend(v int) {
	n := &node{} // Le & crée un pointeur
	n.data = v
	l.append(n)
}

// ============================================
// SUPPRIMER UN NOEUD EN FONCTION DE SA VALEUR
// ============================================

func (l *linkedList) deleteNodeByValue(v int) {
	// CAS 1 - LISTE VIDE
	if l.head == nil {
		return
	}
	// Le 1er noeud a-t-il la valeur v ?
	if l.head.data == v {
		// Si oui, je le retire et "head" devient le second noeud
		l.head = l.head.next
		l.length--
	} else {
		// Sinon...
		previous := l.head     // Je définis que "previous" vaut le premier noeud
		current := l.head.next // "current" vaut le second noeud

		for current != nil {
			if current.data == v {
				// Si le noeud actuel a pour data la valeur que je cherche...
				previous.next = current.next // Je fais en sorte que "previous.next" saute un noeud
				current = current.next       // Je change le noeud actuel pour mettre le suivant à la place
				l.length--
			} else {
				// Sinon je continue de parcourir la liste
				previous = current
				current = current.next
			}
		}
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 49}
	node2 := &node{data: 18}
	node4 := &node{data: 11}
	node5 := &node{data: 7}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.intPrepend(16)

	myList.append(node4)
	myList.append(node5)
	myList.intAppend(2)

	myList.printListData()

	// ======
	// TESTS
	// ======

	// Trouver le premier et le dernier noeud
	fmt.Println("Valeur du premier noeud : ", myList.getFirstNode().data)
	fmt.Println("Valeur du dernier noeud : ", myList.getLastNode().data)

	// Trouver le noeud qui a la valeur 49
	fmt.Println("Le noeud ayant la valeur 49 : ", myList.findNodeByValue(49))
	// Vérifier ce qui se passe quand on cherche une valeur inexistante dans la liste
	fmt.Println("Chercher un noeud avec une valeur inexistante : ", myList.findNodeByValue(3))

	// Supprimer le noeud qui a la valeur 49
	myList.deleteNodeByValue(49)
	fmt.Println("Ma liste en ayant retiré 49 :")
	myList.printListData()

	// Vérifier ce qui se passe quand on supprime une valeur inexistante
	myList.deleteNodeByValue(3)
	fmt.Println("Ma liste en ayant retiré une valeur inexistante :")
	myList.printListData()

	// Cas particuliers : si le noeud recherché est en 1ère ou dernière place
	myList.deleteNodeByValue(16)
	myList.deleteNodeByValue(2)
	fmt.Println("Ma liste en ayant retiré le 1er et dernier noeud (par leurs valeurs) :")
	myList.printListData()

	emptyList := linkedList{}
	emptyList.deleteNodeByValue(2)
}
