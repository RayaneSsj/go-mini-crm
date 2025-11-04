package storage

import (
	"errors"
	"strings"
)

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (m *MemoryStore) Add(nom, email string) (*Contact, error) {
	if strings.TrimSpace(nom) == "" {
		return nil, errors.New("le nom ne peut pas etre vide")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("l'email ne peut pas etre vide")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("l'email doit contenir un @")
	}

	contact := &Contact{
		ID:    m.nextID,
		Nom:   strings.TrimSpace(nom),
		Email: strings.TrimSpace(email),
	}

	m.contacts[m.nextID] = contact
	m.nextID++

	return contact, nil
}

func (m *MemoryStore) List() []*Contact {
	contacts := make([]*Contact, 0, len(m.contacts))
	for _, contact := range m.contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (m *MemoryStore) Get(id int) (*Contact, error) {
	contact, exists := m.contacts[id]
	if !exists {
		return nil, ErrContactNotFound
	}
	return contact, nil
}

func (m *MemoryStore) Update(id int, nom, email string) error {
	contact, exists := m.contacts[id]
	if !exists {
		return ErrContactNotFound
	}

	if nom != "" {
		nomTrimmed := strings.TrimSpace(nom)
		if nomTrimmed == "" {
			return errors.New("le nom ne peut pas etre vide")
		}
		contact.Nom = nomTrimmed
	}

	if email != "" {
		emailTrimmed := strings.TrimSpace(email)
		if emailTrimmed == "" {
			return errors.New("l'email ne peut pas etre vide")
		}
		if !strings.Contains(emailTrimmed, "@") {
			return errors.New("l'email doit contenir un @")
		}
		contact.Email = emailTrimmed
	}

	return nil
}

func (m *MemoryStore) Delete(id int) error {
	if _, exists := m.contacts[id]; !exists {
		return ErrContactNotFound
	}
	delete(m.contacts, id)
	return nil
}
