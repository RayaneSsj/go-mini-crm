package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

var contacts = make(map[int]Contact)
var nextID = 1
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	idFlag := flag.Int("id", 0, "ID du contact")
	nomFlag := flag.String("nom", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if len(os.Args) > 1 && nomFlag != nil && *nomFlag != "" && emailFlag != nil && *emailFlag != "" {
		id := 0
		if idFlag != nil {
			id = *idFlag
		}
		nom := *nomFlag
		email := *emailFlag

		if id == 0 {
			id = nextID
			nextID++
		} else {
			if _, exists := contacts[id]; exists {
				fmt.Printf("Erreur: Un contact avec l'ID %d existe déjà\n", id)
				return
			}
			if id >= nextID {
				nextID = id + 1
			}
		}

		contact := Contact{
			ID:    id,
			Nom:   nom,
			Email: email,
		}
		contacts[id] = contact
		fmt.Printf("Contact ajouté avec l'ID: %d\n", id)
		return
	}

	for {
		fmt.Println("\n=== Mini CRM ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Nom: ")
			scanner.Scan()
			nom := scanner.Text()

			fmt.Print("Email: ")
			scanner.Scan()
			email := scanner.Text()

			contact := Contact{
				ID:    nextID,
				Nom:   nom,
				Email: email,
			}
			contacts[nextID] = contact
			fmt.Printf("Contact ajouté avec l'ID: %d\n", nextID)
			nextID++

		case "2":
			if len(contacts) == 0 {
				fmt.Println("Aucun contact trouvé.")
			} else {
				fmt.Println("\n=== Liste des contacts ===")
				for _, contact := range contacts {
					fmt.Printf("ID: %d | Nom: %s | Email: %s\n", contact.ID, contact.Nom, contact.Email)
				}
			}

		case "3":
			fmt.Print("ID du contact à supprimer: ")
			scanner.Scan()
			idStr := scanner.Text()

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID invalide. Veuillez entrer un nombre.")
			} else {
				if _, exists := contacts[id]; exists {
					delete(contacts, id)
					fmt.Printf("Contact avec l'ID %d supprimé.\n", id)
				} else {
					fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", id)
				}
			}

		case "4":
			fmt.Print("ID du contact à mettre à jour: ")
			scanner.Scan()
			idStr := scanner.Text()

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID invalide. Veuillez entrer un nombre.")
			} else {
				contact, exists := contacts[id]
				if !exists {
					fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", id)
				} else {
					fmt.Printf("Nom actuel: %s\n", contact.Nom)
					fmt.Print("Nouveau nom (laisser vide pour ne pas changer): ")
					scanner.Scan()
					newNom := strings.TrimSpace(scanner.Text())
					if newNom != "" {
						contact.Nom = newNom
					}

					fmt.Printf("Email actuel: %s\n", contact.Email)
					fmt.Print("Nouvel email (laisser vide pour ne pas changer): ")
					scanner.Scan()
					newEmail := strings.TrimSpace(scanner.Text())
					if newEmail != "" {
						contact.Email = newEmail
					}

					contacts[id] = contact
					fmt.Printf("Contact avec l'ID %d mis à jour.\n", id)
				}
			}

		case "5":
			fmt.Println("Au revoir!")
			return

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
