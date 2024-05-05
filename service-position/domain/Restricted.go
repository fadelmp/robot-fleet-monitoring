package domain

type Restricted struct {
	Id          string `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:VARCHAR(255);notNull" json:"name"`
	Description string `gorm:"type:VARCHAR(255)" json:"description"`
	Area        []Area
	Base
}
