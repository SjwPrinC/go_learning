package model

import "time"

//映射user 表
type User struct {
	User_Id         int64
	User_Name       string
	Mobile          string
	Password        string
	Email           string
	Create_Time     time.Time
	Last_Visit_Time time.Time
	State           int8
	Is_Vip          bool
}
