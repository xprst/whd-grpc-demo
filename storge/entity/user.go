package entity

import "time"

// User 用户信息
type User struct {
	ID        int32  `xorm:"pk autoincr 'id'"`
	Name      string `xorm:"varchar(255)"`
	Age       int32
	Sex       bool
	CreatedAt time.Time `xorm:"created" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `xorm:"updated" time_format:"2006-01-02 15:04:05"`
}
