package models

type Url struct {
	ID   uint   `gorm:"primary_key"`
	URL  string `gorm:"not null"`
	Code string `gorm:"unique_index;not null"`
}
