package domain

type Robot struct {
	Id        string  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string  `gorm:"type:VARCHAR(255);notNull" json:"name"`
	Longitude float64 `gorm:"type:NUMERIC;notNull" json:"longitude"`
	Latitude  float64 `gorm:"type:NUMERIC;notNull" json:"latitude"`
	Base
}
