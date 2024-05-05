package domain

type Area struct {
	Id         string  `gorm:"primaryKey;autoIncrement" json:"id"`
	Longitude  float64 `gorm:"type:NUMERIC;notNull" json:"longitude"`
	Latitude   float64 `gorm:"type:NUMERIC;notNull" json:"latitude"`
	Restricted Restricted
	Base
}
