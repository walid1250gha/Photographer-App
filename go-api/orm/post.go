package orm

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	User_ID string
	Detail  string
}
