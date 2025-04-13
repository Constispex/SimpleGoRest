package models

import (
	"database/sql/driver"
	"errors"
)

type Category struct {
	CategoryID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	imageIDName string `gorm:"not null"`
}

func (c Category) Value() (driver.Value, error) {
	if c.Name == "" {
		return nil, errors.New("invalid category name")
	}
	return c.Name, nil
}

// Scan implementiert die Scanner-Schnittstelle.
// Es wird verwendet, um einen Datenbankwert in den benutzerdefinierten Typ zu konvertieren.
func (c *Category) Scan(value interface{}) error {
	name, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Category: value is not a string")
	}
	c.Name = name
	return nil
}
