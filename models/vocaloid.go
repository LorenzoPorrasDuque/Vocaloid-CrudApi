package models

type Vocaloid struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Age   int    `json:"age" gorm:"not null"`
	Songs []Song `json:"songs" gorm:"foreignKey:VocaloidID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" `
}
