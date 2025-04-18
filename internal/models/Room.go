package models

import (
	"database/sql/driver"
	"errors"
	_ "gorm.io/gorm"
	_ "time"
)

// User definiert das Datenmodell für einen Benutzer.

type Room struct {
	RoomID      uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	ImageIDName string `gorm:"not null"`
	Frozen      bool   `gorm:"not null"`
}

func (r Room) Value() (driver.Value, error) {
	if r.Name == "" {
		return nil, errors.New("invalid room name")
	}
	return r.Name, nil
}

func (r *Room) Scan(value interface{}) error {
	name, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Room: value is not a string")
	}
	r.Name = name
	return nil
}
