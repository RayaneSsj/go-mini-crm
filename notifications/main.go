package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Send(message string) error {
	if e.Email == "" {
		return errors.New("email vide")
	}
	if !strings.Contains(e.Email, "@") {
		return errors.New("email invalide")
	}
	fmt.Printf("[EMAIL] Envoi a %s: %s\n", e.Email, message)
	return nil
}

type SmsNotifier struct {
	Phone string
}

func (s SmsNotifier) Send(message string) error {
	if s.Phone == "" {
		return errors.New("numero vide")
	}
	if len(s.Phone) < 10 {
		return errors.New("numero trop court")
	}
	fmt.Printf("[SMS] Envoi au %s: %s\n", s.Phone, message)
	return nil
}

type PushNotifier struct {
	DeviceID string
}

func (p PushNotifier) Send(message string) error {
	if p.DeviceID == "" {
		return errors.New("device ID vide")
	}
	fmt.Printf("[PUSH] Notification au device %s: %s\n", p.DeviceID, message)
	return nil
}

var notifiers []Notifier
var scanner = bufio.NewScanner(os.Stdin)

func addNotifier() {
	fmt.Println("\nType de notificateur:")
	fmt.Println("1. Email")
	fmt.Println("2. SMS")
	fmt.Println("3. Push")
	fmt.Print("Votre choix: ")

	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		fmt.Print("Adresse email: ")
		scanner.Scan()
		email := scanner.Text()
		notifiers = append(notifiers, EmailNotifier{Email: email})
		fmt.Println("Email notifier ajoute")

	case "2":
		fmt.Print("Numero de telephone: ")
		scanner.Scan()
		phone := scanner.Text()
		notifiers = append(notifiers, SmsNotifier{Phone: phone})
		fmt.Println("SMS notifier ajoute")

	case "3":
		fmt.Print("Device ID: ")
		scanner.Scan()
		deviceID := scanner.Text()
		notifiers = append(notifiers, PushNotifier{DeviceID: deviceID})
		fmt.Println("Push notifier ajoute")

	default:
		fmt.Println("Choix invalide")
	}
}

func listNotifiers() {
	if len(notifiers) == 0 {
		fmt.Println("\nAucun notificateur configure")
		return
	}

	fmt.Println("\n=== Liste des notificateurs ===")
	for i, notifier := range notifiers {
		switch n := notifier.(type) {
		case EmailNotifier:
			fmt.Printf("%d. Email: %s\n", i+1, n.Email)
		case SmsNotifier:
			fmt.Printf("%d. SMS: %s\n", i+1, n.Phone)
		case PushNotifier:
			fmt.Printf("%d. Push: %s\n", i+1, n.DeviceID)
		}
	}
}

func sendNotification() {
	if len(notifiers) == 0 {
		fmt.Println("\nAucun notificateur configure. Ajoutez-en d'abord.")
		return
	}

	fmt.Print("\nMessage a envoyer: ")
	scanner.Scan()
	message := scanner.Text()

	if message == "" {
		fmt.Println("Message vide, annulation")
		return
	}

	fmt.Println("\n--- Envoi en cours ---")
	successCount := 0
	errorCount := 0

	for i, notifier := range notifiers {
		err := notifier.Send(message)
		if err != nil {
			fmt.Printf("  [ERREUR %d] %v\n", i+1, err)
			errorCount++
		} else {
			successCount++
		}
	}

	fmt.Printf("\nResultat: %d succes, %d erreurs\n", successCount, errorCount)
}

func removeNotifier() {
	if len(notifiers) == 0 {
		fmt.Println("\nAucun notificateur a supprimer")
		return
	}

	listNotifiers()
	fmt.Print("\nNumero du notificateur a supprimer: ")
	scanner.Scan()
	input := scanner.Text()

	var index int
	fmt.Sscanf(input, "%d", &index)
	index--

	if index < 0 || index >= len(notifiers) {
		fmt.Println("Numero invalide")
		return
	}

	notifiers = append(notifiers[:index], notifiers[index+1:]...)
	fmt.Println("Notificateur supprime")
}

func main() {
	for {
		fmt.Println("\n=== Systeme de Notifications ===")
		fmt.Println("1. Ajouter un notificateur")
		fmt.Println("2. Lister les notificateurs")
		fmt.Println("3. Envoyer une notification")
		fmt.Println("4. Supprimer un notificateur")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addNotifier()
		case "2":
			listNotifiers()
		case "3":
			sendNotification()
		case "4":
			removeNotifier()
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}
