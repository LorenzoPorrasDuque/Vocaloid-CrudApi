package models

type Vocaloid struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Age   int    `json:"age" gorm:"not null"`
	Songs []Song `json:"songs" gorm:"foreignKey:VocaloidID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" `
}
type Song struct {
	SongID      int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" `
	Description string `json:"description"`
	VocaloidID  int    `json:"vocaloid_id" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
