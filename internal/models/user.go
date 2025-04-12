package models

// User definiert das Datenmodell für einen Benutzer.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
