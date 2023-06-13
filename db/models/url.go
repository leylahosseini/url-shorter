package models

type Url struct {
	ID   uint   `json:"primary_key"`
	URL  string `json:"url"`
	Code string `json:"unique_index;not null"`
}

// type Url struct {
// 	//	gorm.Model
// 	ID   uint   `json:"primary_key"`
// 	URL  string `json:"not null"`
// 	Code string `json:"unique_index;not null"`
// }
