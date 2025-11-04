package app

import (
	"bufio"
	"fmt"
	"go-mini-crm/internal/storage"
	"os"
	"strconv"
)

func HandleAdd(store storage.Storer, scanner *bufio.Scanner) {
	fmt.Print("Nom: ")
	scanner.Scan()
	nom := scanner.Text()

	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()

	contact, err := store.Add(nom, email)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}
	fmt.Printf("Contact ajoute avec l'ID: %d\n", contact.ID)
}

func HandleList(store storage.Storer) {
	contacts := store.List()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouve.")
		return
	}

	fmt.Println("\n=== Liste des contacts ===")
	for _, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", contact.ID, contact.Nom, contact.Email)
	}
}

func HandleDelete(store storage.Storer, scanner *bufio.Scanner) {
	fmt.Print("ID du contact a supprimer: ")
	scanner.Scan()
	idStr := scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide. Veuillez entrer un nombre.")
		return
	}

	err = store.Delete(id)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}
	fmt.Printf("Contact avec l'ID %d supprime.\n", id)
}

func HandleUpdate(store storage.Storer, scanner *bufio.Scanner) {
	fmt.Print("ID du contact a mettre a jour: ")
	scanner.Scan()
	idStr := scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide. Veuillez entrer un nombre.")
		return
	}

	contact, err := store.Get(id)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}

	fmt.Printf("Nom actuel: %s\n", contact.Nom)
	fmt.Print("Nouveau nom (laisser vide pour ne pas changer): ")
	scanner.Scan()
	newNom := scanner.Text()

	fmt.Printf("Email actuel: %s\n", contact.Email)
	fmt.Print("Nouvel email (laisser vide pour ne pas changer): ")
	scanner.Scan()
	newEmail := scanner.Text()

	err = store.Update(id, newNom, newEmail)
	if err != nil {
		fmt.Printf("Erreur lors de la mise a jour: %v\n", err)
		return
	}
	fmt.Printf("Contact avec l'ID %d mis a jour.\n", id)
}

func Run(store storage.Storer) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Mini CRM ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre a jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			HandleAdd(store, scanner)
		case "2":
			HandleList(store)
		case "3":
			HandleDelete(store, scanner)
		case "4":
			HandleUpdate(store, scanner)
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide, veuillez reessayer.")
		}
	}
}
