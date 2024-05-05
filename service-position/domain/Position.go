package domain

type Position struct {
	Id             string  `gorm:"primaryKey;autoIncrement" json:"id"`
	RobotId        string  `gorm:"type:VARCHAR(255);notNull" json:"robot_id"`
	RobotName      string  `gorm:"type:VARCHAR(255);notNull" json:"robot_name"`
	RestrictedId   string  `gorm:"type:VARCHAR(255)" json:"restricted_id"`
	RestrictedName string  `gorm:"type:VARCHAR(255)" json:"restricted_name"`
	Longitude      float64 `gorm:"type:NUMERIC;notNull" json:"longitude"`
	Latitude       float64 `gorm:"type:NUMERIC;notNull" json:"latitude"`
	Base
}
