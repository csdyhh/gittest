package models

import "time"

type Employee struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username""`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IDNumber   string    `json:"IDNumber"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser int       `json:"createUser"`
	UpdateUser int       `json:"updateUser"`
}

func CreateEmployee(employee *Employee) error {
	return DB.Create(employee).Error

}
