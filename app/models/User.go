package models

import "fmt"

type UserModel struct {
	UserId int `gorm:"column:id" uri:"id" binding:"required,gt=0"`
	UserName string `gorm:"column:user_name"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) String() string {
	return fmt.Sprintf("user_id:%d,user_name:%s",this.UserId,this.UserName)
}
