package entities

import "database/sql"

type Client struct {
	Idx      uint64       `gorm:"column:idx"`
	Code     string       `gorm:"column:code"`
	IsAdmin  sql.NullBool `gorm:"column:is_admin"`
	Email    string       `gorm:"column:email"`
	Password string       `gorm:"column:password"`
}

func (e *Client) TableName() string {
	return "client"
}
