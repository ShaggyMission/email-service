package models

type User struct {
    ID        string `gorm:"primaryKey;column:id"`
    FirstName string `gorm:"column:firstName"`
    LastName  string `gorm:"column:lastName"`
    Email     string `gorm:"column:email;unique"`
    Password  string `gorm:"column:password"`
    Phone     string `gorm:"column:phone"`
}

func (User) TableName() string {
    return "Users"
}
