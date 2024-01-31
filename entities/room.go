package entities

import "time"

type Room struct {
	Idx          int64     `gorm:"column:idx;primaryKey;autoIncrement:true" json:"idx"`
	Name         string    `gorm:"column:name" json:"name"`
	RegisterDate time.Time `gorm:"column:register_date" json:"register_date"`
}

func (r *Room) TableName() string {
	return "room"
}
