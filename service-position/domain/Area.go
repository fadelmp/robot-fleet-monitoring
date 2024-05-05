package domain

type Area struct {
	Id           string  `gorm:"primaryKey;autoIncrement" json:"id"`
	RestrictedId string  `gorm:"type:VARCHAR(255);notNull" json:"restricted_id"`
	Longitude    float64 `gorm:"type:NUMERIC;notNull" json:"longitude"`
	Latitude     float64 `gorm:"type:NUMERIC;notNull" json:"latitude"`
	Base
}
