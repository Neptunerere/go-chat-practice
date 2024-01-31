package entities

import "database/sql"

type Employee struct {
	Idx      uint64       `gorm:"column:idx"`
	Code     string       `gorm:"column:code"`
	IsAdmin  sql.NullBool `gorm:"column:is_admin"`
	Email    string       `gorm:"column:email"`
	Password string       `gorm:"column:password"`
}

func (e *Employee) TableName() string {
	return "employee"
}
