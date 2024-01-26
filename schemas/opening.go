package schemas

import (
	"gorm.io/gorm"
)


type Opening struct {
	gorm.Model

	Role string
	Company string
	Remote bool
	Link string
	Salary int64

}