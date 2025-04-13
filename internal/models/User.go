package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex;not null"`
	Items      []Item `gorm:"foreignKey:UserID"`
	Containers []Room `gorm:"foreignKey:UserID"`
}
