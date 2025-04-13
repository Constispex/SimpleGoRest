package models

type Item struct {
	ItemID          uint `gorm:"primaryKey"`
	UserID          uint
	Name            string
	Note            string
	ImageIDName     string
	PhotoFilePath   string
	CategoryID      uint     // Fremdschlüssel
	Category        Category `gorm:"foreignKey:CategoryID;references:ID"`
	Status          string
	DaysTillExpired int
	RoomID          uint
}
