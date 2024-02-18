package models

type Song struct {
	SongID      int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" `
	Description string `json:"description"`
	VocaloidID  int    `json:"vocaloid_id" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
