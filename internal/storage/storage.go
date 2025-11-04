package storage

import "errors"

type Contact struct {
	ID    int
	Nom   string
	Email string
}

var ErrContactNotFound = errors.New("contact non trouve")
var ErrContactExists = errors.New("contact existe deja")

type Storer interface {
	Add(nom, email string) (*Contact, error)
	List() []*Contact
	Get(id int) (*Contact, error)
	Update(id int, nom, email string) error
	Delete(id int) error
}
